/*
 * Copyright (C) 2020-2022, IrineSistiana
 *
 * This file is part of mosdns.
 *
 * mosdns is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * mosdns is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package sequence

// import (
// 	"reflect"
// 	"testing"
// )

// func Test_parseExec(t *testing.T) {

// 	tests := []struct {
// 		name string
// 		args string
// 		want ExecConfig
// 	}{
// 		{"", " $t1   a 1  ", ExecConfig{
// 			Tag:     "",
// 			Type:    "typ",
// 			Args:    "a 1",
// 		}},
// 		{"", " typ   a 1  ", ExecConfig{
// 			Tag:     "",
// 			Type:    "typ",
// 			Args:    "a 1",
// 		}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := parseExec(tt.args); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("parseExec() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_parseMatch(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		args string
// 		want MatchConfig
// 	}{
// 		{"", " $m1  a 1 ", MatchConfig{
// 			Tag:     "m1",
// 			Type:    "",
// 			Args:    "a 1",
// 			Reverse: false,
// 		}},
// 		{"", " ! typ  a 1 ", MatchConfig{
// 			Tag:     "",
// 			Type:    "typ",
// 			Args:    "a 1",
// 			Reverse: true,
// 		}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := parseMatch(tt.args); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("parseMatch() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
