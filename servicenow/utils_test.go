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

func TestStringNotEqual(t *testing.T) {
	quals, cols := makeQualMapSingle("category", proto.ColumnType_STRING, "<>",
		&proto.QualValue{Value: &proto.QualValue_StringValue{StringValue: "software"}})
	result := buildQueryFromQuals(quals, cols, nil)
	if result != "category!=software" {
		t.Errorf("expected 'category!=software', got '%s'", result)
	}
}

func TestStringListValueSkipped(t *testing.T) {
	cols := []*plugin.Column{{Name: "category", Type: proto.ColumnType_STRING}}
	quals := plugin.KeyColumnQualMap{
		"category": &plugin.KeyColumnQuals{
			Name: "category",
			Quals: quals.QualSlice{
				&quals.Qual{
					Column:   "category",
					Operator: "=",
					Value: &proto.QualValue{Value: &proto.QualValue_ListValue{
						ListValue: &proto.QualValueList{Values: []*proto.QualValue{
							{Value: &proto.QualValue_StringValue{StringValue: "software"}},
							{Value: &proto.QualValue_StringValue{StringValue: "hardware"}},
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

func TestRowMatchesQualsListQualIgnored(t *testing.T) {
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
	// Lists go to the API with ServiceNow's IN operator, which is exact, so no recheck here
	row := map[string]interface{}{"priority": "3"}
	if !rowMatchesQuals(row, quals, cols) {
		t.Error("expected a list qual to be left to the IN filter")
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
