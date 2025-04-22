package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c:d:e:f", ":")
	want := []string{"a", "b", "c", "d", "e", "f"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestMoreSplit(t *testing.T) {
	got := Split("henryhelloworldhellogolanghellocczuhellobocloud", "hello")
	want := []string{"henry", "world", "golang", "cczu", "bocloud"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSplitChinese(t *testing.T) {
	type Task struct {
		str  string
		sep  string
		want []string
	}

	tasks := map[string]Task{
		"end":   {str: "我好你好大家好", sep: "好", want: []string{"我", "你", "大家"}},
		"start": {str: "一心一意一针一线一天一夜", sep: "一", want: []string{"心", "意", "针", "线", "天", "夜"}},
	}

	for name, task := range tasks {
		t.Run(name, func(t *testing.T) {
			got := Split(task.str, task.sep)
			want := task.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %#v want %#v", got, want)
			}
		})
	}
}
