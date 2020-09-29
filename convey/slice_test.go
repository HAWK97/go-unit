package convey

import (
	"testing"
)

func TestSuccessStringInSlice(t *testing.T) {
	if ok := StringInSlice("a", []string{"a", "b"}); ok {
		t.Log("pass")
	} else {
		t.Error("faild")
	}
}

func TestFaildStringInSlice(t *testing.T) {
	if ok := StringInSlice("b", []string{"a", "b"}); ok {
		//t.Error("faild")
		t.Log("pass")
	} else {
		//t.Log("pass")
		t.Error("faild")
	}
}

//func TestIntegerStuff(t *testing.T) {
//	Convey("Given some integer with a starting value", t, func() {
//		x := 1
//
//		Convey("When the integer is incremented", func() {
//			x++
//
//			Convey("The value should be greater by one", func() {
//				So(x, ShouldEqual, 2)
//			})
//		})
//	})
//}

//func TestStringSliceEqual(t *testing.T) {
//	Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
//		a := []string{"hello", "goconvey"}
//		b := []string{"hello", "goconvey"}
//		So(StringSliceEqual(a, b), ShouldBeTrue)
//	})
//
//	Convey("TestStringSliceEqual should return true when a ＝= nil  && b ＝= nil", t, func() {
//		So(StringSliceEqual(nil, nil), ShouldBeTrue)
//	})
//
//	Convey("TestStringSliceEqual should return false when a ＝= nil  && b != nil", t, func() {
//		a := []string(nil)
//		b := []string{}
//		So(StringSliceEqual(a, b), ShouldBeFalse)
//	})
//
//	Convey("TestStringSliceEqual should return false when a != nil  && b != nil", t, func() {
//		a := []string{"hello", "world"}
//		b := []string{"hello", "goconvey"}
//		So(StringSliceEqual(a, b), ShouldBeFalse)
//	})
//
//}
