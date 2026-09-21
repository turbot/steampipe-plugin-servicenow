package servicenow

import (
	"strings"
	"testing"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/quals"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func makeQualMapSingle(name string, colType proto.ColumnType, op string, val *proto.QualValue) (plugin.KeyColumnQualMap, []*plugin.Column) {
	cols := []*plugin.Column{{Name: name, Type: colType}}
	qm := plugin.KeyColumnQualMap{
		name: &plugin.KeyColumnQuals{
			Name: name,
			Quals: quals.QualSlice{
				&quals.Qual{
					Column:   name,
					Operator: op,
					Value:    val,
				},
			},
		},
	}
	return qm, cols
}

func TestStringEquality(t *testing.T) {
	quals, cols := makeQualMapSingle("category", proto.ColumnType_STRING, "=",
		&proto.QualValue{Value: &proto.QualValue_StringValue{StringValue: "software"}})
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "category=software" {
		t.Errorf("expected 'category=software', got '%s'", result)
	}
}

func TestStringNotEqualNotPushedDown(t *testing.T) {
	// ServiceNow's != is the complement of a case-insensitive match, so it excludes rows that
	// differ only by case and that Postgres would keep. Nothing client-side can add them back,
	// so the filter stays client-side entirely.
	quals, cols := makeQualMapSingle("category", proto.ColumnType_STRING, "<>",
		&proto.QualValue{Value: &proto.QualValue_StringValue{StringValue: "software"}})
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "" {
		t.Errorf("expected empty for a string <>, got '%s'", result)
	}
}

// makeStringListQual builds a list qual on a STRING column, the shape an IN clause produces.
func makeStringListQual(name string, op string, values ...string) (plugin.KeyColumnQualMap, *plugin.Column) {
	listValues := make([]*proto.QualValue, len(values))
	for i, v := range values {
		listValues[i] = &proto.QualValue{Value: &proto.QualValue_StringValue{StringValue: v}}
	}
	qm := plugin.KeyColumnQualMap{
		name: &plugin.KeyColumnQuals{
			Name: name,
			Quals: quals.QualSlice{
				&quals.Qual{
					Column:   name,
					Operator: op,
					Value: &proto.QualValue{Value: &proto.QualValue_ListValue{
						ListValue: &proto.QualValueList{Values: listValues},
					}},
				},
			},
		},
	}
	return qm, &plugin.Column{Name: name, Type: proto.ColumnType_STRING}
}

func TestStringListPushedDownAsIn(t *testing.T) {
	quals, col := makeStringListQual("category", "=", "software", "hardware")
	result := buildQueryFromQuals(quals, []*plugin.Column{col}, nil)
	if result != "categoryINsoftware,hardware" {
		t.Errorf("expected 'categoryINsoftware,hardware', got '%s'", result)
	}
}

func TestStringListSingleValuePushedDownAsIn(t *testing.T) {
	quals, col := makeStringListQual("category", "=", "software")
	result := buildQueryFromQuals(quals, []*plugin.Column{col}, nil)
	if result != "categoryINsoftware" {
		t.Errorf("expected 'categoryINsoftware', got '%s'", result)
	}
}

func TestStringListWithSpacesPushedDownAsIn(t *testing.T) {
	// Spaces need no handling: the client URL encodes them and IN matches the whole value.
	quals, col := makeStringListQual("category", "=", "Event Management", "hardware")
	result := buildQueryFromQuals(quals, []*plugin.Column{col}, nil)
	if result != "categoryINEvent Management,hardware" {
		t.Errorf("expected 'categoryINEvent Management,hardware', got '%s'", result)
	}
}

func TestStringNotInListNotPushedDown(t *testing.T) {
	quals, col := makeStringListQual("category", "<>", "software", "hardware")
	result := buildQueryFromQuals(quals, []*plugin.Column{col}, nil)
	if result != "" {
		t.Errorf("expected empty for a <> list, got '%s'", result)
	}
}

func TestStringEmptyListSkipped(t *testing.T) {
	quals, col := makeStringListQual("category", "=")
	result := buildQueryFromQuals(quals, []*plugin.Column{col}, nil)
	if result != "" {
		t.Errorf("expected empty for an empty list, got '%s'", result)
	}
}

func TestTwoStringListsPushedDownAsIn(t *testing.T) {
	// Two lists are passed through by the SDK unaltered, and the FDW still pushes the limit
	// down for them, so both have to reach the API.
	qm, categoryCol := makeStringListQual("category", "=", "hardware", "network")
	classQuals, classCol := makeStringListQual("sys_class_name", "=", "incident", "problem")
	qm["sys_class_name"] = classQuals["sys_class_name"]
	result := buildQueryFromQuals(qm, []*plugin.Column{categoryCol, classCol}, nil)
	expected := "categoryINhardware,network^sys_class_nameINincident,problem"
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

// An encoded query has no escaping: ^ separates conditions and , separates IN values. The API
// silently drops a malformed fragment instead of rejecting it, so a value carrying a separator
// is not pushed down at all and the unfiltered rows are narrowed client-side.

func TestStringEqualityWithCaretNotPushedDown(t *testing.T) {
	quals, cols := makeQualMapSingle("category", proto.ColumnType_STRING, "=",
		&proto.QualValue{Value: &proto.QualValue_StringValue{StringValue: "hardware^active=true"}})
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "" {
		t.Errorf("expected empty for a value containing ^, got '%s'", result)
	}
}

func TestStringListWithCaretNotPushedDown(t *testing.T) {
	quals, col := makeStringListQual("category", "=", "software", "hard^ware")
	result := buildQueryFromQuals(quals, []*plugin.Column{col}, nil)
	if result != "" {
		t.Errorf("expected empty for a list value containing ^, got '%s'", result)
	}
}

func TestStringListWithCommaNotPushedDown(t *testing.T) {
	// Pushing only the safe values would narrow the filter, so the whole list is dropped.
	quals, col := makeStringListQual("category", "=", "software", "hard,ware")
	result := buildQueryFromQuals(quals, []*plugin.Column{col}, nil)
	if result != "" {
		t.Errorf("expected empty for a list value containing a comma, got '%s'", result)
	}
}

func TestStringListWithCaretStillAllowsOtherFilters(t *testing.T) {
	// Dropping one unpushable list must not drop the filters around it.
	qm, categoryCol := makeStringListQual("category", "=", "hard^ware")
	qm["priority"] = &plugin.KeyColumnQuals{
		Name: "priority",
		Quals: quals.QualSlice{
			&quals.Qual{Column: "priority", Operator: "=",
				Value: &proto.QualValue{Value: &proto.QualValue_Int64Value{Int64Value: 1}}},
		},
	}
	cols := []*plugin.Column{categoryCol, {Name: "priority", Type: proto.ColumnType_INT}}
	result := buildQueryFromQuals(qm, cols, nil)
	if result != "priority=1" {
		t.Errorf("expected 'priority=1', got '%s'", result)
	}
}

func TestIntEquality(t *testing.T) {
	quals, cols := makeQualMapSingle("priority", proto.ColumnType_INT, "=",
		&proto.QualValue{Value: &proto.QualValue_Int64Value{Int64Value: 1}})
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "priority=1" {
		t.Errorf("expected 'priority=1', got '%s'", result)
	}
}

func TestIntGreaterThan(t *testing.T) {
	quals, cols := makeQualMapSingle("priority", proto.ColumnType_INT, ">",
		&proto.QualValue{Value: &proto.QualValue_Int64Value{Int64Value: 2}})
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "priority>2" {
		t.Errorf("expected 'priority>2', got '%s'", result)
	}
}

func TestIntListPushedDownAsIn(t *testing.T) {
	cols := []*plugin.Column{{Name: "priority", Type: proto.ColumnType_INT}}
	quals := plugin.KeyColumnQualMap{
		"priority": &plugin.KeyColumnQuals{
			Name: "priority",
			Quals: quals.QualSlice{
				&quals.Qual{
					Column:   "priority",
					Operator: "=",
					Value: &proto.QualValue{Value: &proto.QualValue_ListValue{
						ListValue: &proto.QualValueList{Values: []*proto.QualValue{
							{Value: &proto.QualValue_Int64Value{Int64Value: 1}},
							{Value: &proto.QualValue_Int64Value{Int64Value: 2}},
						}},
					}},
				},
			},
		},
	}
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "priorityIN1,2" {
		t.Errorf("expected 'priorityIN1,2', got '%s'", result)
	}
}

func TestTwoIntListsPushedDownAsIn(t *testing.T) {
	cols := []*plugin.Column{
		{Name: "priority", Type: proto.ColumnType_INT},
		{Name: "state", Type: proto.ColumnType_INT},
	}
	quals := plugin.KeyColumnQualMap{
		"priority": &plugin.KeyColumnQuals{
			Name: "priority",
			Quals: quals.QualSlice{
				&quals.Qual{
					Column:   "priority",
					Operator: "=",
					Value: &proto.QualValue{Value: &proto.QualValue_ListValue{
						ListValue: &proto.QualValueList{Values: []*proto.QualValue{
							{Value: &proto.QualValue_Int64Value{Int64Value: 1}},
							{Value: &proto.QualValue_Int64Value{Int64Value: 2}},
						}},
					}},
				},
			},
		},
		"state": &plugin.KeyColumnQuals{
			Name: "state",
			Quals: quals.QualSlice{
				&quals.Qual{
					Column:   "state",
					Operator: "=",
					Value: &proto.QualValue{Value: &proto.QualValue_ListValue{
						ListValue: &proto.QualValueList{Values: []*proto.QualValue{
							{Value: &proto.QualValue_Int64Value{Int64Value: 1}},
							{Value: &proto.QualValue_Int64Value{Int64Value: 2}},
						}},
					}},
				},
			},
		},
	}
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "priorityIN1,2^stateIN1,2" {
		t.Errorf("expected 'priorityIN1,2^stateIN1,2', got '%s'", result)
	}
}

// makeIntListQual builds a list qual on an INT column, the shape an IN clause produces.
func makeIntListQual(name string, op string, values ...int64) (plugin.KeyColumnQualMap, *plugin.Column) {
	listValues := make([]*proto.QualValue, len(values))
	for i, v := range values {
		listValues[i] = &proto.QualValue{Value: &proto.QualValue_Int64Value{Int64Value: v}}
	}
	qm := plugin.KeyColumnQualMap{
		name: &plugin.KeyColumnQuals{
			Name: name,
			Quals: quals.QualSlice{
				&quals.Qual{
					Column:   name,
					Operator: op,
					Value: &proto.QualValue{Value: &proto.QualValue_ListValue{
						ListValue: &proto.QualValueList{Values: listValues},
					}},
				},
			},
		},
	}
	return qm, &plugin.Column{Name: name, Type: proto.ColumnType_INT}
}

func TestIntListSingleValuePushedDownAsIn(t *testing.T) {
	quals, col := makeIntListQual("priority", "=", 1)
	result := buildQueryFromQuals(quals, []*plugin.Column{col}, nil)
	if result != "priorityIN1" {
		t.Errorf("expected 'priorityIN1', got '%s'", result)
	}
}

func TestIntEmptyListSkipped(t *testing.T) {
	quals, col := makeIntListQual("priority", "=")
	result := buildQueryFromQuals(quals, []*plugin.Column{col}, nil)
	if result != "" {
		t.Errorf("expected empty for an empty list, got '%s'", result)
	}
}

func TestIntNotInListNotPushedDown(t *testing.T) {
	// A `not in` list arrives with the <> operator. ServiceNow's IN would invert the
	// meaning, so only `=` lists are pushed down and this must stay client-side.
	quals, col := makeIntListQual("priority", "<>", 1, 2)
	result := buildQueryFromQuals(quals, []*plugin.Column{col}, nil)
	if result != "" {
		t.Errorf("expected empty for a <> list, got '%s'", result)
	}
}

func TestIntListWithTimestampQual(t *testing.T) {
	// The shape of a bounded two-list query: the IN filter and the widened date bound
	// both reach the API in one sysparm_query.
	ts := time.Date(2026, 8, 15, 0, 0, 0, 0, time.UTC)
	qm, priorityCol := makeIntListQual("priority", "=", 1, 2)
	qm["opened_at"] = &plugin.KeyColumnQuals{
		Name: "opened_at",
		Quals: quals.QualSlice{
			&quals.Qual{Column: "opened_at", Operator: ">=",
				Value: &proto.QualValue{Value: &proto.QualValue_TimestampValue{TimestampValue: timestamppb.New(ts)}}},
		},
	}
	cols := []*plugin.Column{priorityCol, {Name: "opened_at", Type: proto.ColumnType_TIMESTAMP}}
	result := buildQueryFromQuals(qm, cols, nil)
	expected := "priorityIN1,2^opened_at>=2026-08-14 10:00:00"
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestBoolEquality(t *testing.T) {
	quals, cols := makeQualMapSingle("active", proto.ColumnType_BOOL, "=",
		&proto.QualValue{Value: &proto.QualValue_BoolValue{BoolValue: false}})
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "active=false" {
		t.Errorf("expected 'active=false', got '%s'", result)
	}
}

func TestBoolListValueSkipped(t *testing.T) {
	cols := []*plugin.Column{{Name: "active", Type: proto.ColumnType_BOOL}}
	quals := plugin.KeyColumnQualMap{
		"active": &plugin.KeyColumnQuals{
			Name: "active",
			Quals: quals.QualSlice{
				&quals.Qual{
					Column:   "active",
					Operator: "=",
					Value: &proto.QualValue{Value: &proto.QualValue_ListValue{
						ListValue: &proto.QualValueList{Values: []*proto.QualValue{
							{Value: &proto.QualValue_BoolValue{BoolValue: true}},
						}},
					}},
				},
			},
		},
	}
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "" {
		t.Errorf("expected empty (list skipped), got '%s'", result)
	}
}

func TestTimestampGreaterThanWidened(t *testing.T) {
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	quals, cols := makeQualMapSingle("opened_at", proto.ColumnType_TIMESTAMP, ">",
		&proto.QualValue{Value: &proto.QualValue_TimestampValue{TimestampValue: timestamppb.New(ts)}})
	result := buildQueryFromQuals(quals, cols, nil)
	// > should widen to >= with -14h offset
	expected := "opened_at>=2023-12-31 10:00:00"
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestTimestampLessThanWidened(t *testing.T) {
	ts := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)
	quals, cols := makeQualMapSingle("opened_at", proto.ColumnType_TIMESTAMP, "<",
		&proto.QualValue{Value: &proto.QualValue_TimestampValue{TimestampValue: timestamppb.New(ts)}})
	result := buildQueryFromQuals(quals, cols, nil)
	// < should widen to <= with +14h offset
	expected := "opened_at<=2024-02-01 14:00:00"
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestTimestampEqualWidened(t *testing.T) {
	ts := time.Date(2024, 6, 15, 12, 30, 0, 0, time.UTC)
	quals, cols := makeQualMapSingle("resolved_at", proto.ColumnType_TIMESTAMP, "=",
		&proto.QualValue{Value: &proto.QualValue_TimestampValue{TimestampValue: timestamppb.New(ts)}})
	result := buildQueryFromQuals(quals, cols, nil)
	// = should become a 28h window
	if !strings.Contains(result, "resolved_at>=2024-06-14 22:30:00") {
		t.Errorf("expected lower bound 2024-06-14 22:30:00 in '%s'", result)
	}
	if !strings.Contains(result, "resolved_at<=2024-06-16 02:30:00") {
		t.Errorf("expected upper bound 2024-06-16 02:30:00 in '%s'", result)
	}
	if !strings.Contains(result, "^") {
		t.Errorf("expected ^ separator in '%s'", result)
	}
}

func TestTimestampRangeBothBounds(t *testing.T) {
	ts1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	ts2 := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)
	cols := []*plugin.Column{{Name: "resolved_at", Type: proto.ColumnType_TIMESTAMP}}
	quals := plugin.KeyColumnQualMap{
		"resolved_at": &plugin.KeyColumnQuals{
			Name: "resolved_at",
			Quals: quals.QualSlice{
				&quals.Qual{Column: "resolved_at", Operator: ">=",
					Value: &proto.QualValue{Value: &proto.QualValue_TimestampValue{TimestampValue: timestamppb.New(ts1)}}},
				&quals.Qual{Column: "resolved_at", Operator: "<",
					Value: &proto.QualValue{Value: &proto.QualValue_TimestampValue{TimestampValue: timestamppb.New(ts2)}}},
			},
		},
	}
	result := buildQueryFromQuals(quals, cols, nil)
	// >= 2024-01-01 widened to >= 2023-12-31 10:00:00
	// < 2024-02-01 widened to <= 2024-02-01 14:00:00
	if !strings.Contains(result, "resolved_at>=2023-12-31 10:00:00") {
		t.Errorf("expected widened lower bound in '%s'", result)
	}
	if !strings.Contains(result, "resolved_at<=2024-02-01 14:00:00") {
		t.Errorf("expected widened upper bound in '%s'", result)
	}
}

func TestTimestampListValueSkipped(t *testing.T) {
	cols := []*plugin.Column{{Name: "opened_at", Type: proto.ColumnType_TIMESTAMP}}
	quals := plugin.KeyColumnQualMap{
		"opened_at": &plugin.KeyColumnQuals{
			Name: "opened_at",
			Quals: quals.QualSlice{
				&quals.Qual{
					Column:   "opened_at",
					Operator: "=",
					Value: &proto.QualValue{Value: &proto.QualValue_ListValue{
						ListValue: &proto.QualValueList{Values: []*proto.QualValue{
							{Value: &proto.QualValue_TimestampValue{TimestampValue: timestamppb.New(time.Now())}},
						}},
					}},
				},
			},
		},
	}
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "" {
		t.Errorf("expected empty (list skipped), got '%s'", result)
	}
}

func TestSnowOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"=", "="},
		{"<>", "!="},
		{">", ">"},
		{">=", ">="},
		{"<", "<"},
		{"<=", "<="},
		{"LIKE", ""},
		{"~~", ""},
	}
	for _, tt := range tests {
		result := snowOperator(tt.input)
		if result != tt.expected {
			t.Errorf("snowOperator(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestBoolEqualityTrue(t *testing.T) {
	quals, cols := makeQualMapSingle("active", proto.ColumnType_BOOL, "=",
		&proto.QualValue{Value: &proto.QualValue_BoolValue{BoolValue: true}})
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "active=true" {
		t.Errorf("expected 'active=true', got '%s'", result)
	}
}

func TestBoolNotEqualTrue(t *testing.T) {
	quals, cols := makeQualMapSingle("active", proto.ColumnType_BOOL, "<>",
		&proto.QualValue{Value: &proto.QualValue_BoolValue{BoolValue: true}})
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "active!=true" {
		t.Errorf("expected 'active!=true', got '%s'", result)
	}
}

func TestBoolNotEqualFalse(t *testing.T) {
	quals, cols := makeQualMapSingle("active", proto.ColumnType_BOOL, "<>",
		&proto.QualValue{Value: &proto.QualValue_BoolValue{BoolValue: false}})
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "active!=false" {
		t.Errorf("expected 'active!=false', got '%s'", result)
	}
}

func TestDoubleEquality(t *testing.T) {
	quals, cols := makeQualMapSingle("cost", proto.ColumnType_DOUBLE, "=",
		&proto.QualValue{Value: &proto.QualValue_DoubleValue{DoubleValue: 99.95}})
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "cost=99.95" {
		t.Errorf("expected 'cost=99.95', got '%s'", result)
	}
}

func TestDoubleGreaterThan(t *testing.T) {
	quals, cols := makeQualMapSingle("cost", proto.ColumnType_DOUBLE, ">",
		&proto.QualValue{Value: &proto.QualValue_DoubleValue{DoubleValue: 100.5}})
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "cost>100.5" {
		t.Errorf("expected 'cost>100.5', got '%s'", result)
	}
}

func TestDoubleHighPrecision(t *testing.T) {
	quals, cols := makeQualMapSingle("rate", proto.ColumnType_DOUBLE, ">=",
		&proto.QualValue{Value: &proto.QualValue_DoubleValue{DoubleValue: 0.0000006}})
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "rate>=0.0000006" {
		t.Errorf("expected 'rate>=0.0000006', got '%s'", result)
	}
}

func TestDoubleListValueSkipped(t *testing.T) {
	cols := []*plugin.Column{{Name: "cost", Type: proto.ColumnType_DOUBLE}}
	quals := plugin.KeyColumnQualMap{
		"cost": &plugin.KeyColumnQuals{
			Name: "cost",
			Quals: quals.QualSlice{
				&quals.Qual{
					Column:   "cost",
					Operator: "=",
					Value: &proto.QualValue{Value: &proto.QualValue_ListValue{
						ListValue: &proto.QualValueList{Values: []*proto.QualValue{
							{Value: &proto.QualValue_DoubleValue{DoubleValue: 1.5}},
							{Value: &proto.QualValue_DoubleValue{DoubleValue: 2.5}},
						}},
					}},
				},
			},
		},
	}
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "" {
		t.Errorf("expected empty (list skipped), got '%s'", result)
	}
}

func TestMultiColumnScalarQuals(t *testing.T) {
	cols := []*plugin.Column{
		{Name: "category", Type: proto.ColumnType_STRING},
		{Name: "priority", Type: proto.ColumnType_INT},
	}
	qm := plugin.KeyColumnQualMap{
		"category": &plugin.KeyColumnQuals{
			Name: "category",
			Quals: quals.QualSlice{
				&quals.Qual{Column: "category", Operator: "=",
					Value: &proto.QualValue{Value: &proto.QualValue_StringValue{StringValue: "software"}}},
			},
		},
		"priority": &plugin.KeyColumnQuals{
			Name: "priority",
			Quals: quals.QualSlice{
				&quals.Qual{Column: "priority", Operator: "=",
					Value: &proto.QualValue{Value: &proto.QualValue_Int64Value{Int64Value: 1}}},
			},
		},
	}
	result := buildQueryFromQuals(qm, cols, nil)
	if !strings.Contains(result, "category=software") {
		t.Errorf("expected 'category=software' in '%s'", result)
	}
	if !strings.Contains(result, "priority=1") {
		t.Errorf("expected 'priority=1' in '%s'", result)
	}
	if !strings.Contains(result, "^") {
		t.Errorf("expected '^' separator in '%s'", result)
	}
}

func TestIntLessThanOrEqual(t *testing.T) {
	quals, cols := makeQualMapSingle("priority", proto.ColumnType_INT, "<=",
		&proto.QualValue{Value: &proto.QualValue_Int64Value{Int64Value: 3}})
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "priority<=3" {
		t.Errorf("expected 'priority<=3', got '%s'", result)
	}
}

func TestNilValueSkipped(t *testing.T) {
	cols := []*plugin.Column{{Name: "priority", Type: proto.ColumnType_INT}}
	quals := plugin.KeyColumnQualMap{
		"priority": &plugin.KeyColumnQuals{
			Name: "priority",
			Quals: quals.QualSlice{
				&quals.Qual{Column: "priority", Operator: "=", Value: nil},
			},
		},
	}
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "" {
		t.Errorf("expected empty for nil value, got '%s'", result)
	}
}

//// rowMatchesQuals

// makeTimestampWindow builds the quals for `name >= lower and name < upper`, the shape a
// bounded date range produces. The API call for it is widened by 14h on each side, so
// rowMatchesQuals is what keeps rows outside the real window off the limit.
func makeTimestampWindow(name string, lower, upper time.Time) (plugin.KeyColumnQualMap, []*plugin.Column) {
	cols := []*plugin.Column{{Name: name, Type: proto.ColumnType_TIMESTAMP}}
	qm := plugin.KeyColumnQualMap{
		name: &plugin.KeyColumnQuals{
			Name: name,
			Quals: quals.QualSlice{
				&quals.Qual{Column: name, Operator: ">=",
					Value: &proto.QualValue{Value: &proto.QualValue_TimestampValue{TimestampValue: timestamppb.New(lower)}}},
				&quals.Qual{Column: name, Operator: "<",
					Value: &proto.QualValue{Value: &proto.QualValue_TimestampValue{TimestampValue: timestamppb.New(upper)}}},
			},
		},
	}
	return qm, cols
}

var (
	windowLower = time.Date(2026, 8, 15, 0, 0, 0, 0, time.UTC)
	windowUpper = time.Date(2026, 8, 16, 0, 0, 0, 0, time.UTC)
)

func TestRowMatchesQualsInsideWidenedBoundButOutsideWindow(t *testing.T) {
	quals, cols := makeTimestampWindow("opened_at", windowLower, windowUpper)
	// Returned by the widened >= 2026-08-14 10:00:00 bound, but outside the real day
	row := map[string]interface{}{"opened_at": "2026-08-14 20:00:00"}
	if rowMatchesQuals(row, quals, cols) {
		t.Error("expected row before the window to be skipped")
	}
}

func TestRowMatchesQualsInsideWindow(t *testing.T) {
	quals, cols := makeTimestampWindow("opened_at", windowLower, windowUpper)
	row := map[string]interface{}{"opened_at": "2026-08-15 12:00:00"}
	if !rowMatchesQuals(row, quals, cols) {
		t.Error("expected row inside the window to match")
	}
}

func TestRowMatchesQualsLowerEdgeInclusive(t *testing.T) {
	quals, cols := makeTimestampWindow("opened_at", windowLower, windowUpper)
	row := map[string]interface{}{"opened_at": "2026-08-15 00:00:00"}
	if !rowMatchesQuals(row, quals, cols) {
		t.Error("expected the >= lower edge to match")
	}
}

func TestRowMatchesQualsUpperEdgeExclusive(t *testing.T) {
	quals, cols := makeTimestampWindow("opened_at", windowLower, windowUpper)
	row := map[string]interface{}{"opened_at": "2026-08-16 00:00:00"}
	if rowMatchesQuals(row, quals, cols) {
		t.Error("expected the < upper edge to be skipped")
	}
}

func TestRowMatchesQualsEmptyValueWithNotEqualQual(t *testing.T) {
	quals, cols := makeQualMapSingle("active", proto.ColumnType_BOOL, "<>",
		&proto.QualValue{Value: &proto.QualValue_BoolValue{BoolValue: true}})
	// ServiceNow's != also returns empty values, which sanitizeTableObject turns into nil
	row := map[string]interface{}{"active": nil}
	if rowMatchesQuals(row, quals, cols) {
		t.Error("expected an empty value to be skipped for a <> qual")
	}
}

func TestRowMatchesQualsIntQualPassesThrough(t *testing.T) {
	quals, cols := makeQualMapSingle("priority", proto.ColumnType_INT, "=",
		&proto.QualValue{Value: &proto.QualValue_Int64Value{Int64Value: 1}})
	// Integer filters are exact in sysparm_query, so a present value always matches
	row := map[string]interface{}{"priority": "1"}
	if !rowMatchesQuals(row, quals, cols) {
		t.Error("expected an int qual with a present value to match")
	}
}

func TestRowMatchesQualsMissingFieldIgnored(t *testing.T) {
	quals, cols := makeTimestampWindow("opened_at", windowLower, windowUpper)
	row := map[string]interface{}{"number": "INC0010001"}
	if !rowMatchesQuals(row, quals, cols) {
		t.Error("expected a row missing the qualified field to be left alone")
	}
}

func TestRowMatchesQualsIntListMembershipRechecked(t *testing.T) {
	// The IN filter is exact for integers, but a `not in` list is never pushed down at all, so
	// membership is rechecked either way rather than trusting the filter that was sent.
	quals, col := makeIntListQual("priority", "=", 1, 2)
	cols := []*plugin.Column{col}
	if !rowMatchesQuals(map[string]interface{}{"priority": "2"}, quals, cols) {
		t.Error("expected a value in the list to match")
	}
	if rowMatchesQuals(map[string]interface{}{"priority": "3"}, quals, cols) {
		t.Error("expected a value outside the list to be skipped")
	}
}

func TestRowMatchesQualsIntNotInListRechecked(t *testing.T) {
	quals, col := makeIntListQual("priority", "<>", 1, 2)
	cols := []*plugin.Column{col}
	if !rowMatchesQuals(map[string]interface{}{"priority": "3"}, quals, cols) {
		t.Error("expected a value outside a not-in list to match")
	}
	if rowMatchesQuals(map[string]interface{}{"priority": "1"}, quals, cols) {
		t.Error("expected a value in a not-in list to be skipped")
	}
}

// ServiceNow compares strings case-insensitively, for =, != and IN alike. Only = and IN are
// pushed down, both returning a superset, and these are what narrow that superset exactly.

func TestRowMatchesQualsStringEqualityIsCaseSensitive(t *testing.T) {
	quals, cols := makeQualMapSingle("category", proto.ColumnType_STRING, "=",
		&proto.QualValue{Value: &proto.QualValue_StringValue{StringValue: "event management"}})

	if !rowMatchesQuals(map[string]interface{}{"category": "event management"}, quals, cols) {
		t.Error("expected an exact match to match")
	}
	if rowMatchesQuals(map[string]interface{}{"category": "Event Management"}, quals, cols) {
		t.Error("expected a case variant returned by the API to be skipped")
	}
}

func TestRowMatchesQualsStringNotEqualIsCaseSensitive(t *testing.T) {
	quals, cols := makeQualMapSingle("category", proto.ColumnType_STRING, "<>",
		&proto.QualValue{Value: &proto.QualValue_StringValue{StringValue: "event management"}})

	// This is the row ServiceNow's != would have excluded, and Postgres keeps it
	if !rowMatchesQuals(map[string]interface{}{"category": "Event Management"}, quals, cols) {
		t.Error("expected a case variant to satisfy a string <>")
	}
	if rowMatchesQuals(map[string]interface{}{"category": "event management"}, quals, cols) {
		t.Error("expected an exact match to be skipped for a string <>")
	}
}

func TestRowMatchesQualsStringListMembershipIsCaseSensitive(t *testing.T) {
	quals, col := makeStringListQual("category", "=", "hardware", "network")
	cols := []*plugin.Column{col}

	if !rowMatchesQuals(map[string]interface{}{"category": "network"}, quals, cols) {
		t.Error("expected a value in the list to match")
	}
	if rowMatchesQuals(map[string]interface{}{"category": "Network"}, quals, cols) {
		t.Error("expected a case variant of a list value to be skipped")
	}
	if rowMatchesQuals(map[string]interface{}{"category": "software"}, quals, cols) {
		t.Error("expected a value outside the list to be skipped")
	}
}

func TestRowMatchesQualsStringNotInListRechecked(t *testing.T) {
	// A `not in` list is never pushed down, so the API returns everything and this is the only
	// thing keeping non-matching rows from consuming the pushed down limit.
	quals, col := makeStringListQual("category", "<>", "hardware", "network")
	cols := []*plugin.Column{col}

	if !rowMatchesQuals(map[string]interface{}{"category": "software"}, quals, cols) {
		t.Error("expected a value outside a not-in list to match")
	}
	if !rowMatchesQuals(map[string]interface{}{"category": "Hardware"}, quals, cols) {
		t.Error("expected a case variant to satisfy a not-in list")
	}
	if rowMatchesQuals(map[string]interface{}{"category": "hardware"}, quals, cols) {
		t.Error("expected a value in a not-in list to be skipped")
	}
}

func TestRowMatchesQualsEmptyValueWithListQual(t *testing.T) {
	// IN never returns empty values, but an unpushable list leaves them in the response, and
	// NULL satisfies neither `in` nor `not in` in Postgres.
	inQuals, col := makeStringListQual("category", "=", "hardware")
	cols := []*plugin.Column{col}
	if rowMatchesQuals(map[string]interface{}{"category": nil}, inQuals, cols) {
		t.Error("expected an empty value to be skipped for an in list")
	}
	notInQuals, _ := makeStringListQual("category", "<>", "hardware")
	if rowMatchesQuals(map[string]interface{}{"category": nil}, notInQuals, cols) {
		t.Error("expected an empty value to be skipped for a not-in list")
	}
}

func TestRowMatchesQualsBoolListMembership(t *testing.T) {
	// ServiceNow returns booleans as strings
	cols := []*plugin.Column{{Name: "active", Type: proto.ColumnType_BOOL}}
	qm := plugin.KeyColumnQualMap{
		"active": &plugin.KeyColumnQuals{
			Name: "active",
			Quals: quals.QualSlice{
				&quals.Qual{Column: "active", Operator: "=",
					Value: &proto.QualValue{Value: &proto.QualValue_ListValue{
						ListValue: &proto.QualValueList{Values: []*proto.QualValue{
							{Value: &proto.QualValue_BoolValue{BoolValue: true}},
						}},
					}},
				},
			},
		},
	}
	if !rowMatchesQuals(map[string]interface{}{"active": "true"}, qm, cols) {
		t.Error("expected true to match a list of (true)")
	}
	if rowMatchesQuals(map[string]interface{}{"active": "false"}, qm, cols) {
		t.Error("expected false to be skipped for a list of (true)")
	}
}

func TestRowMatchesQualsTimestampListMembership(t *testing.T) {
	// Timestamp lists are not pushed down, so the whole table comes back and membership has to
	// be applied here to the instant, not the string.
	cols := []*plugin.Column{{Name: "opened_at", Type: proto.ColumnType_TIMESTAMP}}
	wanted := time.Date(2026, 8, 15, 12, 0, 0, 0, time.UTC)
	qm := plugin.KeyColumnQualMap{
		"opened_at": &plugin.KeyColumnQuals{
			Name: "opened_at",
			Quals: quals.QualSlice{
				&quals.Qual{Column: "opened_at", Operator: "=",
					Value: &proto.QualValue{Value: &proto.QualValue_ListValue{
						ListValue: &proto.QualValueList{Values: []*proto.QualValue{
							{Value: &proto.QualValue_TimestampValue{TimestampValue: timestamppb.New(wanted)}},
						}},
					}},
				},
			},
		},
	}
	if !rowMatchesQuals(map[string]interface{}{"opened_at": "2026-08-15 12:00:00"}, qm, cols) {
		t.Error("expected the listed instant to match")
	}
	if rowMatchesQuals(map[string]interface{}{"opened_at": "2026-08-15 12:00:01"}, qm, cols) {
		t.Error("expected an instant one second off to be skipped")
	}
}

func TestRowMatchesQualsDoubleListComparesNumerically(t *testing.T) {
	// ServiceNow's string form of a number need not match Postgres's, so 1.50 is 1.5
	cols := []*plugin.Column{{Name: "cost", Type: proto.ColumnType_DOUBLE}}
	qm := plugin.KeyColumnQualMap{
		"cost": &plugin.KeyColumnQuals{
			Name: "cost",
			Quals: quals.QualSlice{
				&quals.Qual{Column: "cost", Operator: "=",
					Value: &proto.QualValue{Value: &proto.QualValue_ListValue{
						ListValue: &proto.QualValueList{Values: []*proto.QualValue{
							{Value: &proto.QualValue_DoubleValue{DoubleValue: 1.5}},
						}},
					}},
				},
			},
		},
	}
	if !rowMatchesQuals(map[string]interface{}{"cost": "1.50"}, qm, cols) {
		t.Error("expected 1.50 to match a list of (1.5)")
	}
	if rowMatchesQuals(map[string]interface{}{"cost": "2.5"}, qm, cols) {
		t.Error("expected 2.5 to be skipped for a list of (1.5)")
	}
}

func TestRowMatchesQualsUnparsableListValueKept(t *testing.T) {
	// A value the plugin cannot read tells it nothing, so the row goes to Postgres rather than
	// being dropped on a guess.
	quals, col := makeIntListQual("priority", "=", 1, 2)
	cols := []*plugin.Column{col}
	if !rowMatchesQuals(map[string]interface{}{"priority": "not a number"}, quals, cols) {
		t.Error("expected an unparsable int to be left to Postgres")
	}
}

// The operators below are widened or relaxed on the way to the API, so rowMatchesQuals is the
// only thing enforcing them exactly before a row counts towards a pushed down limit.

func TestRowMatchesQualsEqualityRequiresExactTimestamp(t *testing.T) {
	// buildQueryFromQuals turns a timestamp = into a 28h window, so the whole window comes back
	ts := time.Date(2026, 8, 15, 12, 0, 0, 0, time.UTC)
	quals, cols := makeQualMapSingle("opened_at", proto.ColumnType_TIMESTAMP, "=",
		&proto.QualValue{Value: &proto.QualValue_TimestampValue{TimestampValue: timestamppb.New(ts)}})

	if !rowMatchesQuals(map[string]interface{}{"opened_at": "2026-08-15 12:00:00"}, quals, cols) {
		t.Error("expected the exact instant to match")
	}
	if rowMatchesQuals(map[string]interface{}{"opened_at": "2026-08-15 12:00:01"}, quals, cols) {
		t.Error("expected an instant one second off to be skipped")
	}
}

func TestRowMatchesQualsStrictGreaterThanExcludesBoundary(t *testing.T) {
	// > is sent to the API as >=, so the boundary row comes back and must be rejected here
	ts := time.Date(2026, 8, 15, 0, 0, 0, 0, time.UTC)
	quals, cols := makeQualMapSingle("opened_at", proto.ColumnType_TIMESTAMP, ">",
		&proto.QualValue{Value: &proto.QualValue_TimestampValue{TimestampValue: timestamppb.New(ts)}})

	if rowMatchesQuals(map[string]interface{}{"opened_at": "2026-08-15 00:00:00"}, quals, cols) {
		t.Error("expected the boundary instant to be skipped for >")
	}
	if !rowMatchesQuals(map[string]interface{}{"opened_at": "2026-08-15 00:00:01"}, quals, cols) {
		t.Error("expected an instant after the boundary to match for >")
	}
}

func TestRowMatchesQualsLessThanOrEqualIncludesBoundary(t *testing.T) {
	ts := time.Date(2026, 8, 15, 0, 0, 0, 0, time.UTC)
	quals, cols := makeQualMapSingle("opened_at", proto.ColumnType_TIMESTAMP, "<=",
		&proto.QualValue{Value: &proto.QualValue_TimestampValue{TimestampValue: timestamppb.New(ts)}})

	if !rowMatchesQuals(map[string]interface{}{"opened_at": "2026-08-15 00:00:00"}, quals, cols) {
		t.Error("expected the boundary instant to match for <=")
	}
	if rowMatchesQuals(map[string]interface{}{"opened_at": "2026-08-15 00:00:01"}, quals, cols) {
		t.Error("expected an instant after the boundary to be skipped for <=")
	}
}

func TestRowMatchesQualsQualZoneIsInstantBased(t *testing.T) {
	// The FDW sends an instant, not a wall clock reading: 2026-08-14T20:00:00-04:00 is the
	// same instant as 2026-08-15T00:00:00Z, which is how the plugin logs it.
	lower := time.Date(2026, 8, 14, 20, 0, 0, 0, time.FixedZone("EDT", -4*60*60))
	quals, cols := makeQualMapSingle("opened_at", proto.ColumnType_TIMESTAMP, ">=",
		&proto.QualValue{Value: &proto.QualValue_TimestampValue{TimestampValue: timestamppb.New(lower)}})

	if !rowMatchesQuals(map[string]interface{}{"opened_at": "2026-08-15 00:00:00"}, quals, cols) {
		t.Error("expected the equivalent UTC instant to match the offset bound")
	}
	if rowMatchesQuals(map[string]interface{}{"opened_at": "2026-08-14 23:59:59"}, quals, cols) {
		t.Error("expected an instant before the offset bound to be skipped")
	}
}

// A row is only skipped when a qual is known not to match. Anything the plugin cannot evaluate
// is left to the Postgres recheck rather than dropped.

func TestRowMatchesQualsUnparsableTimestampKept(t *testing.T) {
	quals, cols := makeTimestampWindow("opened_at", windowLower, windowUpper)
	row := map[string]interface{}{"opened_at": "not a datetime"}
	if !rowMatchesQuals(row, quals, cols) {
		t.Error("expected an unparsable datetime to be left to Postgres")
	}
}

func TestRowMatchesQualsNonStringTimestampKept(t *testing.T) {
	quals, cols := makeTimestampWindow("opened_at", windowLower, windowUpper)
	row := map[string]interface{}{"opened_at": 1755216000}
	if !rowMatchesQuals(row, quals, cols) {
		t.Error("expected a non-string datetime to be left to Postgres")
	}
}

func TestRowMatchesQualsUnsupportedOperatorKept(t *testing.T) {
	// LIKE is never pushed down, so it must not influence the recheck either
	quals, cols := makeQualMapSingle("short_description", proto.ColumnType_STRING, "~~",
		&proto.QualValue{Value: &proto.QualValue_StringValue{StringValue: "disk%"}})
	row := map[string]interface{}{"short_description": "network outage"}
	if !rowMatchesQuals(row, quals, cols) {
		t.Error("expected an operator we never pushed down to be left to Postgres")
	}
}

func TestRowMatchesQualsMultiColumnRejectsOnSecondColumn(t *testing.T) {
	// The timestamp matches, but ServiceNow's != also returned a row with an empty active
	qm, cols := makeTimestampWindow("opened_at", windowLower, windowUpper)
	qm["active"] = &plugin.KeyColumnQuals{
		Name: "active",
		Quals: quals.QualSlice{
			&quals.Qual{Column: "active", Operator: "<>",
				Value: &proto.QualValue{Value: &proto.QualValue_BoolValue{BoolValue: true}}},
		},
	}
	cols = append(cols, &plugin.Column{Name: "active", Type: proto.ColumnType_BOOL})
	row := map[string]interface{}{"opened_at": "2026-08-15 12:00:00", "active": nil}
	if rowMatchesQuals(row, qm, cols) {
		t.Error("expected a row failing the second qual to be skipped")
	}
}

func TestRowMatchesQualsAfterSanitizeTableObject(t *testing.T) {
	// sanitizeTableObject is what turns ServiceNow's empty string into the nil that the
	// recheck rejects, so the two have to stay in step
	quals, cols := makeQualMapSingle("active", proto.ColumnType_BOOL, "<>",
		&proto.QualValue{Value: &proto.QualValue_BoolValue{BoolValue: true}})
	row := map[string]interface{}{"active": ""}
	if !rowMatchesQuals(row, quals, cols) {
		t.Error("expected the raw empty string to pass before sanitizing")
	}
	sanitizeTableObject(row)
	if rowMatchesQuals(row, quals, cols) {
		t.Error("expected the sanitized empty value to be skipped")
	}
}

func TestRowMatchesQualsStringRangeLeftToPostgres(t *testing.T) {
	// No table declares an ordered operator on a string column today. If one ever does, Postgres
	// collates strings its own way, so a range must not be second-guessed here - and in
	// particular the value on the bound must not be dropped as if this were a <>.
	for _, operator := range []string{">", ">=", "<", "<="} {
		quals, cols := makeQualMapSingle("category", proto.ColumnType_STRING, operator,
			&proto.QualValue{Value: &proto.QualValue_StringValue{StringValue: "hardware"}})
		if !rowMatchesQuals(map[string]interface{}{"category": "hardware"}, quals, cols) {
			t.Errorf("expected the bound value to be left to Postgres for %s", operator)
		}
		if !rowMatchesQuals(map[string]interface{}{"category": "software"}, quals, cols) {
			t.Errorf("expected a non-matching value to be left to Postgres for %s", operator)
		}
	}
}

func TestRowMatchesQualsOrderedListLeftToPostgres(t *testing.T) {
	// A list only ever arrives with = or <>. Anything else is not a membership test, so it must
	// not be read as one.
	quals, col := makeIntListQual("priority", ">", 1, 2)
	if !rowMatchesQuals(map[string]interface{}{"priority": "1"}, quals, []*plugin.Column{col}) {
		t.Error("expected an ordered list operator to be left to Postgres")
	}
}
