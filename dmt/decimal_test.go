package dmt

import (
	"fmt"
	"log/slog"
	"math"
	"testing"

	"github.com/cockroachdb/apd/v3"
)

func TestAddTo(t *testing.T) {
	type args struct {
		x *apd.Decimal
		y *apd.Decimal
		z *apd.Decimal
	}
	tests := []struct {
		args args
		want *apd.Decimal
	}{
		{
			args: args{
				x: apd.New(1234, -3),
				y: apd.New(1234, -3),
				z: new(apd.Decimal),
			},
			want: apd.New(2468, -3),
		},
		{
			args: args{
				x: apd.New(1234, -3),
				y: apd.New(1234, -3),
				z: new(apd.Decimal),
			},
			want: apd.New(2468, -3),
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("testcase_%d", i), func(t *testing.T) {
			actual, err := AddTo(tt.args.z, tt.args.x, tt.args.y)
			if err != nil {
				t.Fatal(err)
			}
			slog.Info(actual.Text('f'))
			if actual.Cmp(tt.want) != 0 {
				t.Errorf("AddTo() = %v, want %v", actual, tt.want)
			}
		})
	}
}

func TestMulTo(t *testing.T) {
	x, _, _ := apd.NewFromString("3.14159265358979323846264338327950288419716939937510582097494459")
	y := apd.New(2134, -3)
	z := new(apd.Decimal)
	actual, err := MulTo(z, x, y)
	if err != nil {
		t.Fatal(err)
	}
	f64, _ := actual.Float64()
	const expect = math.Pi * 2.134
	diff := math.Abs(f64 - expect)
	slog.Info(actual.Text('f'), "diff", diff)
	if diff > 0.0001 {
		t.Errorf("MulTo() = %v, want %v", actual, expect)
	}
}

func TestIntegralTo(t *testing.T) {
	type args struct {
		x *apd.Decimal
		y *apd.BigInt
	}
	tests := []struct {
		args args
		want *apd.BigInt
	}{
		{
			args: args{
				x: apd.New(1234567, -2),
				y: new(apd.BigInt),
			},
			want: apd.NewBigInt(12345),
		},
		{
			args: args{
				x: apd.New(1234567, 2),
				y: new(apd.BigInt),
			},
			want: apd.NewBigInt(123456700),
		},
		{
			args: args{
				x: apd.New(1234567, -122),
				y: apd.NewBigInt(10),
			},
			want: apd.NewBigInt(0),
		},
		{
			args: args{
				x: apd.New(-1234567, -2),
				y: new(apd.BigInt),
			},
			want: apd.NewBigInt(-12345),
		},
		{
			args: args{
				x: apd.New(-1234567, 2),
				y: new(apd.BigInt),
			},
			want: apd.NewBigInt(-123456700),
		},
		{
			args: args{
				x: apd.New(1234567, 122),
				y: apd.NewBigInt(10),
			},
			want: func() *apd.BigInt {
				integer := apd.NewBigInt(1234567)
				for range 122 {
					integer.Mul(integer, new(apd.BigInt).SetInt64(10))
				}
				return integer
			}(),
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("testcase_%d", i), func(t *testing.T) {
			got := IntegralTo(tt.args.y, tt.args.x)
			fmt.Printf("IntegralTo(%v) => %v\n", tt.args.x, got)
			if got.Cmp(tt.want) != 0 {
				t.Errorf("IntegralTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetScale(t *testing.T) {
	type args struct {
		d         *Decimal
		precision int
		scale     int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				d:         New(12345678, -3),
				precision: 40,
				scale:     6,
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		if tt.name == "" {
			tt.name = "testcase"
		}
		t.Run(fmt.Sprintf("%s_%d", tt.name, i), func(t *testing.T) {
			if err := SetScale(tt.args.d, tt.args.precision, tt.args.scale, apd.RoundHalfUp); (err != nil) != tt.wantErr {
				t.Errorf("SetScale() error = %v, wantErr %v", err, tt.wantErr)
			}
			integ, frac := new(Decimal), new(Decimal)
			tt.args.d.Modf(integ, frac)
			fmt.Println(integ.Coeff.Text(10), integ.Sign(), integ.Text('f'))
			fmt.Println(frac.Coeff.Text(10), frac.Sign(), frac.Text('f'))
		})
	}
}

func BenchmarkIntegral1(b *testing.B) {
	type args struct {
		x *apd.Decimal
		y *apd.BigInt
	}
	type bench struct {
		args args
		want *apd.BigInt
	}
	benches := []bench{
		{
			args: args{
				x: apd.New(1234567, -2),
				y: new(apd.BigInt),
			},
			want: apd.NewBigInt(12345),
		},
		{
			args: args{
				x: apd.New(1234567, 2),
				y: new(apd.BigInt),
			},
			want: apd.NewBigInt(123456700),
		},
		{
			args: args{
				x: apd.New(1234567, -122),
				y: new(apd.BigInt),
			},
			want: apd.NewBigInt(0),
		},
		{
			args: args{
				x: apd.New(-1234567, -2),
				y: new(apd.BigInt),
			},
			want: apd.NewBigInt(-12345),
		},
		{
			args: args{
				x: apd.New(-1234567, 2),
				y: new(apd.BigInt),
			},
			want: apd.NewBigInt(-123456700),
		},
		{
			args: args{
				x: apd.New(1234567, 122),
				y: new(apd.BigInt),
			},
			want: func() *apd.BigInt {
				integer := apd.NewBigInt(1234567)
				for range 122 {
					integer.Mul(integer, new(apd.BigInt).SetInt64(10))
				}
				return integer
			}(),
		},
	}
	b.ResetTimer()
	for range 100000 {
		for _, v := range benches {
			_ = IntegralTo(v.args.y, v.args.x)
		}
	}
}

func BenchmarkIntegral2(b *testing.B) {
	type args struct {
		x *apd.Decimal
		y *apd.BigInt
	}
	type bench struct {
		args args
		want *apd.BigInt
	}
	benches := []bench{
		{
			args: args{
				x: apd.New(1234567, -2),
				y: apd.NewBigInt(10),
			},
			want: apd.NewBigInt(12345),
		},
		{
			args: args{
				x: apd.New(1234567, 2),
				y: apd.NewBigInt(10),
			},
			want: apd.NewBigInt(123456700),
		},
		{
			args: args{
				x: apd.New(1234567, -122),
				y: apd.NewBigInt(10),
			},
			want: apd.NewBigInt(0),
		},
		{
			args: args{
				x: apd.New(-1234567, -2),
				y: apd.NewBigInt(10),
			},
			want: apd.NewBigInt(-12345),
		},
		{
			args: args{
				x: apd.New(-1234567, 2),
				y: apd.NewBigInt(10),
			},
			want: apd.NewBigInt(-123456700),
		},
		{
			args: args{
				x: apd.New(1234567, 122),
				y: apd.NewBigInt(10),
			},
			want: func() *apd.BigInt {
				integer := apd.NewBigInt(1234567)
				for range 122 {
					integer.Mul(integer, new(apd.BigInt).SetInt64(10))
				}
				return integer
			}(),
		},
	}
	b.ResetTimer()
	for range 100000 {
		for _, v := range benches {
			_ = IntegralTo(v.args.y, v.args.x)
		}
	}
}

func TestSafePresision(t *testing.T) {
	type args struct {
		x *Decimal
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				x: New(1234, 3),
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				x: New(-1234, 3),
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				x: New(123, -7),
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				x: New(-123, -7),
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				x: New(1234567, -3),
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				x: New(-1234567, -3),
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				x: New(0, -7),
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				x: New(0, 6),
			},
			want: 7,
		},
	}
	for i, tt := range tests {
		if tt.name == "" {
			tt.name = "testcase"
		}
		t.Run(fmt.Sprintf("%s_%d", tt.name, i), func(t *testing.T) {
			got := SafePresision(tt.args.x)
			if got != tt.want {
				t.Errorf("SafePresision() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumTo(t *testing.T) {
	type args struct {
		dst    *Decimal
		values []*Decimal
	}
	tests := []struct {
		name    string
		args    args
		want    *Decimal
		wantErr bool
	}{
		{
			args: args{
				dst:    &apd.Decimal{},
				values: []*apd.Decimal{New(0, 0), New(1, 0), New(2, 0), New(3, 0), New(4, 0)},
			},
			want:    new(apd.Decimal).SetInt64(10),
			wantErr: false,
		},
		{
			args: args{
				dst:    new(apd.Decimal),
				values: []*apd.Decimal{New(0, 0), New(1, 0), New(2, 0), New(3, 0), New(4, 0)},
			},
			want:    new(apd.Decimal).SetInt64(10),
			wantErr: false,
		},
		{
			args: args{
				dst:    new(Decimal).SetInt64(12),
				values: []*apd.Decimal{New(0, 0), New(1, 0), New(2, 0), New(3, 0), New(4, 0)},
			},
			want:    new(apd.Decimal).SetInt64(10),
			wantErr: false,
		},
	}
	for i, tt := range tests {
		if tt.name == "" {
			tt.name = fmt.Sprintf("testcase_%d", i)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumTo(tt.args.dst, tt.args.values...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Cmp(tt.want) != 0 {
				t.Errorf("SumTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductTo(t *testing.T) {
	type args struct {
		dst    *Decimal
		values []*Decimal
	}
	tests := []struct {
		name    string
		args    args
		want    *Decimal
		wantErr bool
	}{
		{
			args: args{
				dst:    &apd.Decimal{},
				values: []*apd.Decimal{New(0, 0), New(1, 0), New(2, 0), New(3, 0), New(4, 0)},
			},
			want:    new(Decimal),
			wantErr: false,
		},
		{
			args: args{
				dst:    new(Decimal),
				values: []*apd.Decimal{New(1, 0), New(2, 0), New(3, 0), New(4, 0)},
			},
			want:    new(Decimal).SetInt64(24),
			wantErr: false,
		},
		{
			args: args{
				dst:    &apd.Decimal{},
				values: []*apd.Decimal{New(1, 0), New(2, 0), New(3, 0), New(4, 0)},
			},
			want:    new(Decimal).SetInt64(24),
			wantErr: false,
		},
		{
			args: args{
				dst:    new(Decimal).SetInt64(12),
				values: []*apd.Decimal{New(1, 0), New(2, 0), New(3, 0), New(4, 0)},
			},
			want:    new(Decimal).SetInt64(24),
			wantErr: false,
		},
		{
			args: args{
				dst:    new(Decimal).SetInt64(12),
				values: []*apd.Decimal{New(1, 0)},
			},
			want:    new(Decimal).SetInt64(1),
			wantErr: false,
		},
	}
	for i, tt := range tests {
		if tt.name == "" {
			tt.name = fmt.Sprintf("testcase_%d", i)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProductTo(tt.args.dst, tt.args.values...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Cmp(tt.want) != 0 {
				t.Errorf("ProductTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
