package main

import (
	"reflect"
	"testing"
)

func TestNewNumAa(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want numAa
	}{
		{name: "three",
			args: args{3},
			want: numAa{
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
			want: numAa{
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
			want: numAa{
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
			if got := NewNumAa(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNumAa() = %v, want %v", got, tt.want)
			}
		})
	}
	// check length of numAa
	base := len(NewNumAa(0))
	for i := 1; i < 11; i++ {
		if base != len(NewNumAa(i)) {
			t.Errorf("len(NewNumAa(%d) = %d, want %d", i, base, len(NewNumAa(i)))
		}
	}
}

func Test_numAa_join(t *testing.T) {
	tests := []struct {
		name string
		na   numAa
		want string
	}{
		{name: "test one",
			na: NewNumAa(1),
			want: `   #    
  ##    
 # #    
   #    
   #    
   #    
 #####  
`},
		{name: "test five",
			na: NewNumAa(5),
			want: `####### 
#       
#       
######  
      # 
#     # 
 #####  
`},
		{name: "test colon",
			na: NewNumAa(10),
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
				t.Errorf("numAa.join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numAa_merge(t *testing.T) {
	type args struct {
		na1 numAa
	}
	tests := []struct {
		name string
		na   numAa
		args args
		want numAa
	}{
		{name: "one merge two",
			na:   NewNumAa(1),
			args: args{NewNumAa(2)},
			want: numAa{
				{" ", " ", " ", "#", " ", " ", " ", " ", " ", "#", "#", "#", "#", "#", " ", " "},
				{" ", " ", "#", "#", " ", " ", " ", " ", "#", " ", " ", " ", " ", " ", "#", " "},
				{" ", "#", " ", "#", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "#", " "},
				{" ", " ", " ", "#", " ", " ", " ", " ", " ", "#", "#", "#", "#", "#", " ", " "},
				{" ", " ", " ", "#", " ", " ", " ", " ", "#", " ", " ", " ", " ", " ", " ", " "},
				{" ", " ", " ", "#", " ", " ", " ", " ", "#", " ", " ", " ", " ", " ", " ", " "},
				{" ", "#", "#", "#", "#", "#", " ", " ", "#", "#", "#", "#", "#", "#", "#", " "},
			}},
		{name: "five merge seven",
			na:   NewNumAa(5),
			args: args{NewNumAa(7)},
			want: numAa{
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
				t.Errorf("numAa.merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
