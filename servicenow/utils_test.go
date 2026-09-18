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

func TestIntListValueSkipped(t *testing.T) {
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
	if result != "" {
		t.Errorf("expected empty (list skipped), got '%s'", result)
	}
}

func TestTwoIntListsSkipped(t *testing.T) {
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
	if result != "" {
		t.Errorf("expected empty (both lists skipped), got '%s'", result)
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
