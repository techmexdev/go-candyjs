package duktape

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type DuktapeSuite struct {
	ctx *Context
}

var _ = Suite(&DuktapeSuite{})

func (s *DuktapeSuite) SetUpTest(c *C) {
	s.ctx = NewContext()
}

func (s *DuktapeSuite) TestRegisterFunc_String(c *C) {
	var called interface{}
	s.ctx.RegisterFunc("test_in_string", func(s string) {
		called = s
	})

	s.ctx.EvalString("test_in_string('foo')")
	c.Assert(called, Equals, "foo")
}

func (s *DuktapeSuite) TestRegisterFunc_Int(c *C) {
	var ri, ri8, ri16, ri32, ri64 interface{}
	s.ctx.RegisterFunc("test_in_int", func(i int, i8 int8, i16 int16, i32 int32, i64 int64) {
		ri = i
		ri8 = i8
		ri16 = i16
		ri32 = i32
		ri64 = i64
	})

	s.ctx.EvalString("test_in_int(42, 8, 16, 32, 64)")
	c.Assert(ri, Equals, 42)
	c.Assert(ri8, Equals, int8(8))
	c.Assert(ri16, Equals, int16(16))
	c.Assert(ri32, Equals, int32(32))
	c.Assert(ri64, Equals, int64(64))
}

func (s *DuktapeSuite) TestRegisterFunc_Uint(c *C) {
	var ri, ri8, ri16, ri32, ri64 interface{}
	s.ctx.RegisterFunc("test_in_uint", func(i uint, i8 uint8, i16 uint16, i32 uint32, i64 uint64) {
		ri = i
		ri8 = i8
		ri16 = i16
		ri32 = i32
		ri64 = i64
	})

	s.ctx.EvalString("test_in_uint(42, 8, 16, 32, 64)")
	c.Assert(ri, Equals, uint(42))
	c.Assert(ri8, Equals, uint8(8))
	c.Assert(ri16, Equals, uint16(16))
	c.Assert(ri32, Equals, uint32(32))
	c.Assert(ri64, Equals, uint64(64))
}

func (s *DuktapeSuite) TestRegisterFunc_Float(c *C) {
	var called64 interface{}
	var called32 interface{}
	s.ctx.RegisterFunc("test_in_float", func(f64 float64, f32 float32) {
		called64 = f64
		called32 = f32
	})

	s.ctx.EvalString("test_in_float(42, 42)")
	c.Assert(called64, Equals, 42.0)
	c.Assert(called32, Equals, float32(42.0))
}

func (s *DuktapeSuite) TestRegisterFunc_Bool(c *C) {
	var called interface{}
	s.ctx.RegisterFunc("test_in_bool", func(b bool) {
		called = b
	})

	s.ctx.EvalString("test_in_bool(true)")
	c.Assert(called, Equals, true)
}

func (s *DuktapeSuite) TestRegisterFunc_Interface(c *C) {
	var called interface{}
	s.ctx.RegisterFunc("test_in_interface", func(i interface{}) {
		called = i
	})

	s.ctx.EvalString("test_in_interface('qux')")
	c.Assert(called, Equals, "qux")
}

func (s *DuktapeSuite) TestRegisterFunc_Slice(c *C) {
	var called interface{}
	s.ctx.RegisterFunc("test_in_slice", func(s []interface{}) {
		called = s
	})

	s.ctx.EvalString("test_in_slice(['foo', 42])")
	c.Assert(called, DeepEquals, []interface{}{"foo", 42.0})
}

func (s *DuktapeSuite) TestRegisterFunc_Map(c *C) {
	var called interface{}
	s.ctx.RegisterFunc("test_in_map", func(s map[string]interface{}) {
		called = s
	})

	s.ctx.EvalString("test_in_map({foo: 42, qux: 'bar'})")

	c.Assert(called, DeepEquals, map[string]interface{}{"foo": 42.0, "qux": "bar"})
}

func (s *DuktapeSuite) TestRegisterFunc_Nil(c *C) {
	var cm, cs, ci, cst interface{}
	s.ctx.RegisterFunc("test_nil", func(m map[string]interface{}, s []interface{}, i int, st string) {
		cm = m
		cs = s
		ci = i
		cst = st
	})

	s.ctx.EvalString("test_nil(null, null, null, null)")
	c.Assert(cm, DeepEquals, map[string]interface{}(nil))
	c.Assert(cs, DeepEquals, []interface{}(nil))
	c.Assert(ci, DeepEquals, 0)
	c.Assert(cst, DeepEquals, "")
}

func (s *DuktapeSuite) TestRegisterFunc_Optional(c *C) {
	var cm, cs, ci, cst interface{}
	s.ctx.RegisterFunc("test_optional", func(m map[string]interface{}, s []interface{}, i int, st string) {
		cm = m
		cs = s
		ci = i
		cst = st
	})

	s.ctx.EvalString("test_optional()")
	c.Assert(cm, DeepEquals, map[string]interface{}(nil))
	c.Assert(cs, DeepEquals, []interface{}(nil))
	c.Assert(ci, DeepEquals, 0)
	c.Assert(cst, DeepEquals, "")
}

func (s *DuktapeSuite) TestRegisterFunc_Variadic(c *C) {
	var calledA interface{}
	var calledB interface{}
	s.ctx.RegisterFunc("test_in_variadic", func(s string, is ...int) {
		calledA = s
		calledB = is
	})

	s.ctx.EvalString("test_in_variadic('foo', 21, 42)")
	c.Assert(calledA, DeepEquals, "foo")
	c.Assert(calledB, DeepEquals, []int{21, 42})
}

func (s *DuktapeSuite) TearDownTest(c *C) {
	s.ctx.DestroyHeap()
}