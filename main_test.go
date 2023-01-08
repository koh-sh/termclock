package main

import (
	"reflect"
	"testing"
)

func TestNewNumaa(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want numaa
	}{
		{name: "three",
			args: args{3},
			want: numaa{
				{" ", "#", "#", "#", "#", "#", " ", " "},
				{"#", " ", " ", " ", " ", " ", "#", " "},
				{" ", " ", " ", " ", " ", " ", "#", " "},
				{" ", "#", "#", "#", "#", "#", " ", " "},
				{" ", " ", " ", " ", " ", " ", "#", " "},
				{"#", " ", " ", " ", " ", " ", "#", " "},
				{" ", "#", "#", "#", "#", "#", " ", " "},
			}},
		{name: "zero",
			args: args{0},
			want: numaa{
				{" ", " ", "#", "#", "#", " ", " ", " "},
				{" ", "#", " ", " ", " ", "#", " ", " "},
				{"#", " ", " ", " ", " ", " ", "#", " "},
				{"#", " ", " ", " ", " ", " ", "#", " "},
				{"#", " ", " ", " ", " ", " ", "#", " "},
				{" ", "#", " ", " ", " ", "#", " ", " "},
				{" ", " ", "#", "#", "#", " ", " ", " "},
			}},
		{name: "colon",
			args: args{10},
			want: numaa{
				{" ", " ", " ", " ", " "},
				{" ", " ", "#", " ", " "},
				{" ", " ", "#", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", "#", " ", " "},
				{" ", " ", "#", " ", " "},
				{" ", " ", " ", " ", " "},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNumaa(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNumaa() = %v, want %v", got, tt.want)
			}
		})
	}
	// check length of numaa
	base := len(NewNumaa(0))
	for i := 1; i < 11; i++ {
		if base != len(NewNumaa(i)) {
			t.Errorf("len(NewNumaa(%d) = %d, want %d", i, base, len(NewNumaa(i)))
		}
	}
}

func Test_numaa_join(t *testing.T) {
	tests := []struct {
		name string
		na   numaa
		want string
	}{
		{name: "test one",
			na: NewNumaa(1),
			want: `   #    
  ##    
 # #    
   #    
   #    
   #    
 #####  
`},
		{name: "test five",
			na: NewNumaa(5),
			want: `####### 
#       
#       
######  
      # 
#     # 
 #####  
`},
		{name: "test colon",
			na: NewNumaa(10),
			want: `     
  #  
  #  
     
  #  
  #  
     
`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.na.join(); got != tt.want {
				t.Errorf("numaa.join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numaa_merge(t *testing.T) {
	type args struct {
		na1 numaa
	}
	tests := []struct {
		name string
		na   numaa
		args args
		want numaa
	}{
		{name: "one merge two",
			na:   NewNumaa(1),
			args: args{NewNumaa(2)},
			want: numaa{
				{" ", " ", " ", "#", " ", " ", " ", " ", " ", "#", "#", "#", "#", "#", " ", " "},
				{" ", " ", "#", "#", " ", " ", " ", " ", "#", " ", " ", " ", " ", " ", "#", " "},
				{" ", "#", " ", "#", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "#", " "},
				{" ", " ", " ", "#", " ", " ", " ", " ", " ", "#", "#", "#", "#", "#", " ", " "},
				{" ", " ", " ", "#", " ", " ", " ", " ", "#", " ", " ", " ", " ", " ", " ", " "},
				{" ", " ", " ", "#", " ", " ", " ", " ", "#", " ", " ", " ", " ", " ", " ", " "},
				{" ", "#", "#", "#", "#", "#", " ", " ", "#", "#", "#", "#", "#", "#", "#", " "},
			}},
		{name: "five merge seven",
			na:   NewNumaa(5),
			args: args{NewNumaa(7)},
			want: numaa{
				{"#", "#", "#", "#", "#", "#", "#", " ", "#", "#", "#", "#", "#", "#", "#", " "},
				{"#", " ", " ", " ", " ", " ", " ", " ", "#", " ", " ", " ", " ", "#", " ", " "},
				{"#", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "#", " ", " ", " "},
				{"#", "#", "#", "#", "#", "#", " ", " ", " ", " ", " ", "#", " ", " ", " ", " "},
				{" ", " ", " ", " ", " ", " ", "#", " ", " ", " ", "#", " ", " ", " ", " ", " "},
				{"#", " ", " ", " ", " ", " ", "#", " ", " ", " ", "#", " ", " ", " ", " ", " "},
				{" ", "#", "#", "#", "#", "#", " ", " ", " ", " ", "#", " ", " ", " ", " ", " "},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.na.merge(tt.args.na1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numaa.merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
