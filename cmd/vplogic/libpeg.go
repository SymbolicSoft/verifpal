/* SPDX-FileCopyrightText: © 2019-2021 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

// This file is generated automatically from libpeg.peg.
// Do not modify.

package vplogic

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var libpegReserved = []string{
	"attacker", "passive", "active", "principal",
	"knows", "generates", "leaks",
	"phase", "public", "private", "password",
	"confidentiality", "authentication",
	"freshness", "unlinkability", "precondition",
	"ringsign", "ringsignverif",
	"primitive", "pw_hash", "hash", "hkdf",
	"aead_enc", "aead_dec", "enc", "dec",
	"mac", "assert", "sign", "signverif",
	"pke_enc", "pke_dec", "shamir_split",
	"shamir_join", "concat", "split",
	"g", "nil", "unnamed",
}

var libpegUnnamedCounter = 0

func libpegCheckIfReserved(s string) error {
	found := false
	switch {
	case strInSlice(s, libpegReserved):
		found = true
	case strings.HasPrefix(s, "attacker"):
		found = true
	case strings.HasPrefix(s, "unnamed"):
		found = true
	}
	if found {
		return fmt.Errorf("cannot use reserved keyword in Name: %s", s)
	}
	return nil
}

func libpegParseModel(filePath string, verbose bool) (Model, error) {
	fileName := filepath.Base(filePath)
	if len(fileName) > 64 {
		return Model{}, fmt.Errorf("model file name must be 64 characters or less")
	}
	if filepath.Ext(fileName) != ".vp" {
		return Model{}, fmt.Errorf("model file name must have a '.vp' extension")
	}
	if verbose {
		InfoMessage(fmt.Sprintf(
			"Parsing model '%s'...", fileName,
		), "verifpal", false)
	}
	parsed, err := ParseFile(filePath)
	if err != nil {
		return Model{}, err
	}
	m := parsed.(Model)
	m.FileName = fileName
	return m, nil
}

var g = &grammar{
	rules: []*rule{
		{
			name: "Model",
			pos:  position{line: 80, col: 1, offset: 1777},
			expr: &actionExpr{
				pos: position{line: 80, col: 10, offset: 1786},
				run: (*parser).callonModel1,
				expr: &seqExpr{
					pos: position{line: 80, col: 10, offset: 1786},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 80, col: 10, offset: 1786},
							expr: &ruleRefExpr{
								pos:  position{line: 80, col: 10, offset: 1786},
								name: "Comment",
							},
						},
						&labeledExpr{
							pos:   position{line: 80, col: 19, offset: 1795},
							label: "Attacker",
							expr: &zeroOrOneExpr{
								pos: position{line: 80, col: 28, offset: 1804},
								expr: &ruleRefExpr{
									pos:  position{line: 80, col: 28, offset: 1804},
									name: "Attacker",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 80, col: 38, offset: 1814},
							label: "Blocks",
							expr: &zeroOrOneExpr{
								pos: position{line: 80, col: 45, offset: 1821},
								expr: &oneOrMoreExpr{
									pos: position{line: 80, col: 46, offset: 1822},
									expr: &ruleRefExpr{
										pos:  position{line: 80, col: 46, offset: 1822},
										name: "Block",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 80, col: 55, offset: 1831},
							label: "Queries",
							expr: &zeroOrOneExpr{
								pos: position{line: 80, col: 63, offset: 1839},
								expr: &ruleRefExpr{
									pos:  position{line: 80, col: 63, offset: 1839},
									name: "Queries",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 80, col: 72, offset: 1848},
							expr: &ruleRefExpr{
								pos:  position{line: 80, col: 72, offset: 1848},
								name: "Comment",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 80, col: 81, offset: 1857},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Attacker",
			pos:  position{line: 102, col: 1, offset: 2409},
			expr: &actionExpr{
				pos: position{line: 102, col: 13, offset: 2421},
				run: (*parser).callonAttacker1,
				expr: &seqExpr{
					pos: position{line: 102, col: 13, offset: 2421},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 102, col: 13, offset: 2421},
							val:        "attacker",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 102, col: 24, offset: 2432},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 102, col: 26, offset: 2434},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 102, col: 30, offset: 2438},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 102, col: 32, offset: 2440},
							label: "Type",
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 37, offset: 2445},
								name: "AttackerType",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 102, col: 50, offset: 2458},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 102, col: 52, offset: 2460},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 102, col: 56, offset: 2464},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "AttackerType",
			pos:  position{line: 106, col: 1, offset: 2489},
			expr: &actionExpr{
				pos: position{line: 106, col: 17, offset: 2505},
				run: (*parser).callonAttackerType1,
				expr: &choiceExpr{
					pos: position{line: 106, col: 18, offset: 2506},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 106, col: 18, offset: 2506},
							val:        "active",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 106, col: 27, offset: 2515},
							val:        "passive",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Block",
			pos:  position{line: 110, col: 1, offset: 2559},
			expr: &actionExpr{
				pos: position{line: 110, col: 10, offset: 2568},
				run: (*parser).callonBlock1,
				expr: &seqExpr{
					pos: position{line: 110, col: 10, offset: 2568},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 110, col: 10, offset: 2568},
							expr: &ruleRefExpr{
								pos:  position{line: 110, col: 10, offset: 2568},
								name: "Comment",
							},
						},
						&labeledExpr{
							pos:   position{line: 110, col: 19, offset: 2577},
							label: "Block",
							expr: &choiceExpr{
								pos: position{line: 110, col: 26, offset: 2584},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 110, col: 26, offset: 2584},
										name: "Phase",
									},
									&ruleRefExpr{
										pos:  position{line: 110, col: 32, offset: 2590},
										name: "Principal",
									},
									&ruleRefExpr{
										pos:  position{line: 110, col: 42, offset: 2600},
										name: "Message",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 110, col: 51, offset: 2609},
							name: "_",
						},
						&zeroOrMoreExpr{
							pos: position{line: 110, col: 53, offset: 2611},
							expr: &ruleRefExpr{
								pos:  position{line: 110, col: 53, offset: 2611},
								name: "Comment",
							},
						},
					},
				},
			},
		},
		{
			name: "Principal",
			pos:  position{line: 114, col: 1, offset: 2644},
			expr: &actionExpr{
				pos: position{line: 114, col: 14, offset: 2657},
				run: (*parser).callonPrincipal1,
				expr: &seqExpr{
					pos: position{line: 114, col: 14, offset: 2657},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 114, col: 14, offset: 2657},
							val:        "principal",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 114, col: 26, offset: 2669},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 114, col: 28, offset: 2671},
							label: "Name",
							expr: &ruleRefExpr{
								pos:  position{line: 114, col: 33, offset: 2676},
								name: "PrincipalName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 114, col: 47, offset: 2690},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 114, col: 49, offset: 2692},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 114, col: 53, offset: 2696},
							name: "_",
						},
						&zeroOrMoreExpr{
							pos: position{line: 114, col: 55, offset: 2698},
							expr: &ruleRefExpr{
								pos:  position{line: 114, col: 55, offset: 2698},
								name: "Comment",
							},
						},
						&labeledExpr{
							pos:   position{line: 114, col: 64, offset: 2707},
							label: "Expressions",
							expr: &zeroOrMoreExpr{
								pos: position{line: 114, col: 77, offset: 2720},
								expr: &ruleRefExpr{
									pos:  position{line: 114, col: 77, offset: 2720},
									name: "Expression",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 114, col: 90, offset: 2733},
							expr: &ruleRefExpr{
								pos:  position{line: 114, col: 90, offset: 2733},
								name: "Comment",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 114, col: 99, offset: 2742},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 114, col: 101, offset: 2744},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 114, col: 105, offset: 2748},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "PrincipalName",
			pos:  position{line: 127, col: 1, offset: 2989},
			expr: &actionExpr{
				pos: position{line: 127, col: 18, offset: 3006},
				run: (*parser).callonPrincipalName1,
				expr: &labeledExpr{
					pos:   position{line: 127, col: 18, offset: 3006},
					label: "Name",
					expr: &ruleRefExpr{
						pos:  position{line: 127, col: 23, offset: 3011},
						name: "Identifier",
					},
				},
			},
		},
		{
			name: "Qualifier",
			pos:  position{line: 132, col: 1, offset: 3114},
			expr: &actionExpr{
				pos: position{line: 132, col: 14, offset: 3127},
				run: (*parser).callonQualifier1,
				expr: &choiceExpr{
					pos: position{line: 132, col: 15, offset: 3128},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 132, col: 15, offset: 3128},
							val:        "public",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 132, col: 24, offset: 3137},
							val:        "private",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 132, col: 34, offset: 3147},
							val:        "password",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Message",
			pos:  position{line: 136, col: 1, offset: 3192},
			expr: &actionExpr{
				pos: position{line: 136, col: 12, offset: 3203},
				run: (*parser).callonMessage1,
				expr: &seqExpr{
					pos: position{line: 136, col: 12, offset: 3203},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 136, col: 12, offset: 3203},
							label: "Sender",
							expr: &ruleRefExpr{
								pos:  position{line: 136, col: 19, offset: 3210},
								name: "PrincipalName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 136, col: 33, offset: 3224},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 136, col: 35, offset: 3226},
							val:        "->",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 136, col: 40, offset: 3231},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 136, col: 42, offset: 3233},
							label: "Recipient",
							expr: &ruleRefExpr{
								pos:  position{line: 136, col: 52, offset: 3243},
								name: "PrincipalName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 136, col: 66, offset: 3257},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 136, col: 68, offset: 3259},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 136, col: 72, offset: 3263},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 136, col: 74, offset: 3265},
							label: "Constants",
							expr: &ruleRefExpr{
								pos:  position{line: 136, col: 84, offset: 3275},
								name: "MessageConstants",
							},
						},
					},
				},
			},
		},
		{
			name: "MessageConstants",
			pos:  position{line: 147, col: 1, offset: 3464},
			expr: &actionExpr{
				pos: position{line: 147, col: 21, offset: 3484},
				run: (*parser).callonMessageConstants1,
				expr: &labeledExpr{
					pos:   position{line: 147, col: 21, offset: 3484},
					label: "MessageConstants",
					expr: &oneOrMoreExpr{
						pos: position{line: 147, col: 38, offset: 3501},
						expr: &choiceExpr{
							pos: position{line: 147, col: 39, offset: 3502},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 147, col: 39, offset: 3502},
									name: "GuardedConstant",
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 55, offset: 3518},
									name: "Constant",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Expression",
			pos:  position{line: 157, col: 1, offset: 3682},
			expr: &actionExpr{
				pos: position{line: 157, col: 15, offset: 3696},
				run: (*parser).callonExpression1,
				expr: &seqExpr{
					pos: position{line: 157, col: 15, offset: 3696},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 157, col: 15, offset: 3696},
							expr: &ruleRefExpr{
								pos:  position{line: 157, col: 15, offset: 3696},
								name: "Comment",
							},
						},
						&labeledExpr{
							pos:   position{line: 157, col: 24, offset: 3705},
							label: "Expression",
							expr: &choiceExpr{
								pos: position{line: 157, col: 36, offset: 3717},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 157, col: 36, offset: 3717},
										name: "Knows",
									},
									&ruleRefExpr{
										pos:  position{line: 157, col: 42, offset: 3723},
										name: "Generates",
									},
									&ruleRefExpr{
										pos:  position{line: 157, col: 52, offset: 3733},
										name: "Leaks",
									},
									&ruleRefExpr{
										pos:  position{line: 157, col: 58, offset: 3739},
										name: "Assignment",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 157, col: 70, offset: 3751},
							name: "_",
						},
						&zeroOrMoreExpr{
							pos: position{line: 157, col: 72, offset: 3753},
							expr: &ruleRefExpr{
								pos:  position{line: 157, col: 72, offset: 3753},
								name: "Comment",
							},
						},
					},
				},
			},
		},
		{
			name: "Knows",
			pos:  position{line: 161, col: 1, offset: 3791},
			expr: &actionExpr{
				pos: position{line: 161, col: 10, offset: 3800},
				run: (*parser).callonKnows1,
				expr: &seqExpr{
					pos: position{line: 161, col: 10, offset: 3800},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 161, col: 10, offset: 3800},
							val:        "knows",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 18, offset: 3808},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 161, col: 20, offset: 3810},
							label: "Qualifier",
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 30, offset: 3820},
								name: "Qualifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 40, offset: 3830},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 161, col: 42, offset: 3832},
							label: "Constants",
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 52, offset: 3842},
								name: "Constants",
							},
						},
					},
				},
			},
		},
		{
			name: "Generates",
			pos:  position{line: 169, col: 1, offset: 3972},
			expr: &actionExpr{
				pos: position{line: 169, col: 14, offset: 3985},
				run: (*parser).callonGenerates1,
				expr: &seqExpr{
					pos: position{line: 169, col: 14, offset: 3985},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 169, col: 14, offset: 3985},
							val:        "generates",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 169, col: 26, offset: 3997},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 169, col: 28, offset: 3999},
							label: "Constants",
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 38, offset: 4009},
								name: "Constants",
							},
						},
					},
				},
			},
		},
		{
			name: "Leaks",
			pos:  position{line: 177, col: 1, offset: 4127},
			expr: &actionExpr{
				pos: position{line: 177, col: 10, offset: 4136},
				run: (*parser).callonLeaks1,
				expr: &seqExpr{
					pos: position{line: 177, col: 10, offset: 4136},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 177, col: 10, offset: 4136},
							val:        "leaks",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 177, col: 18, offset: 4144},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 177, col: 20, offset: 4146},
							label: "Constants",
							expr: &ruleRefExpr{
								pos:  position{line: 177, col: 30, offset: 4156},
								name: "Constants",
							},
						},
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 185, col: 1, offset: 4270},
			expr: &actionExpr{
				pos: position{line: 185, col: 15, offset: 4284},
				run: (*parser).callonAssignment1,
				expr: &seqExpr{
					pos: position{line: 185, col: 15, offset: 4284},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 185, col: 15, offset: 4284},
							label: "Left",
							expr: &ruleRefExpr{
								pos:  position{line: 185, col: 20, offset: 4289},
								name: "Constants",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 185, col: 30, offset: 4299},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 185, col: 32, offset: 4301},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 185, col: 36, offset: 4305},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 185, col: 38, offset: 4307},
							label: "Right",
							expr: &ruleRefExpr{
								pos:  position{line: 185, col: 44, offset: 4313},
								name: "Value",
							},
						},
					},
				},
			},
		},
		{
			name: "Constant",
			pos:  position{line: 206, col: 1, offset: 4757},
			expr: &actionExpr{
				pos: position{line: 206, col: 13, offset: 4769},
				run: (*parser).callonConstant1,
				expr: &seqExpr{
					pos: position{line: 206, col: 13, offset: 4769},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 206, col: 13, offset: 4769},
							label: "Const",
							expr: &ruleRefExpr{
								pos:  position{line: 206, col: 19, offset: 4775},
								name: "Identifier",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 206, col: 30, offset: 4786},
							expr: &seqExpr{
								pos: position{line: 206, col: 31, offset: 4787},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 206, col: 31, offset: 4787},
										name: "_",
									},
									&litMatcher{
										pos:        position{line: 206, col: 33, offset: 4789},
										val:        ",",
										ignoreCase: false,
									},
									&ruleRefExpr{
										pos:  position{line: 206, col: 37, offset: 4793},
										name: "_",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Constants",
			pos:  position{line: 215, col: 1, offset: 4897},
			expr: &actionExpr{
				pos: position{line: 215, col: 14, offset: 4910},
				run: (*parser).callonConstants1,
				expr: &labeledExpr{
					pos:   position{line: 215, col: 14, offset: 4910},
					label: "Constants",
					expr: &oneOrMoreExpr{
						pos: position{line: 215, col: 24, offset: 4920},
						expr: &ruleRefExpr{
							pos:  position{line: 215, col: 24, offset: 4920},
							name: "Constant",
						},
					},
				},
			},
		},
		{
			name: "Phase",
			pos:  position{line: 227, col: 1, offset: 5163},
			expr: &actionExpr{
				pos: position{line: 227, col: 10, offset: 5172},
				run: (*parser).callonPhase1,
				expr: &seqExpr{
					pos: position{line: 227, col: 10, offset: 5172},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 227, col: 10, offset: 5172},
							val:        "phase",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 227, col: 18, offset: 5180},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 227, col: 20, offset: 5182},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 227, col: 24, offset: 5186},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 227, col: 26, offset: 5188},
							label: "Number",
							expr: &oneOrMoreExpr{
								pos: position{line: 227, col: 33, offset: 5195},
								expr: &charClassMatcher{
									pos:        position{line: 227, col: 33, offset: 5195},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 227, col: 40, offset: 5202},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 227, col: 42, offset: 5204},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 227, col: 46, offset: 5208},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "GuardedConstant",
			pos:  position{line: 240, col: 1, offset: 5430},
			expr: &actionExpr{
				pos: position{line: 240, col: 20, offset: 5449},
				run: (*parser).callonGuardedConstant1,
				expr: &seqExpr{
					pos: position{line: 240, col: 20, offset: 5449},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 240, col: 20, offset: 5449},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 240, col: 24, offset: 5453},
							label: "Guarded",
							expr: &ruleRefExpr{
								pos:  position{line: 240, col: 32, offset: 5461},
								name: "Identifier",
							},
						},
						&litMatcher{
							pos:        position{line: 240, col: 43, offset: 5472},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrOneExpr{
							pos: position{line: 240, col: 47, offset: 5476},
							expr: &seqExpr{
								pos: position{line: 240, col: 48, offset: 5477},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 240, col: 48, offset: 5477},
										name: "_",
									},
									&litMatcher{
										pos:        position{line: 240, col: 50, offset: 5479},
										val:        ",",
										ignoreCase: false,
									},
									&ruleRefExpr{
										pos:  position{line: 240, col: 54, offset: 5483},
										name: "_",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Primitive",
			pos:  position{line: 251, col: 1, offset: 5653},
			expr: &actionExpr{
				pos: position{line: 251, col: 14, offset: 5666},
				run: (*parser).callonPrimitive1,
				expr: &seqExpr{
					pos: position{line: 251, col: 14, offset: 5666},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 251, col: 14, offset: 5666},
							label: "Name",
							expr: &ruleRefExpr{
								pos:  position{line: 251, col: 19, offset: 5671},
								name: "PrimitiveName",
							},
						},
						&litMatcher{
							pos:        position{line: 251, col: 33, offset: 5685},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 251, col: 37, offset: 5689},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 251, col: 39, offset: 5691},
							label: "Arguments",
							expr: &oneOrMoreExpr{
								pos: position{line: 251, col: 49, offset: 5701},
								expr: &ruleRefExpr{
									pos:  position{line: 251, col: 49, offset: 5701},
									name: "Value",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 251, col: 56, offset: 5708},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 251, col: 58, offset: 5710},
							val:        ")",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 251, col: 62, offset: 5714},
							label: "Check",
							expr: &zeroOrOneExpr{
								pos: position{line: 251, col: 68, offset: 5720},
								expr: &litMatcher{
									pos:        position{line: 251, col: 68, offset: 5720},
									val:        "?",
									ignoreCase: false,
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 251, col: 73, offset: 5725},
							expr: &seqExpr{
								pos: position{line: 251, col: 74, offset: 5726},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 251, col: 74, offset: 5726},
										name: "_",
									},
									&litMatcher{
										pos:        position{line: 251, col: 76, offset: 5728},
										val:        ",",
										ignoreCase: false,
									},
									&ruleRefExpr{
										pos:  position{line: 251, col: 80, offset: 5732},
										name: "_",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "PrimitiveName",
			pos:  position{line: 267, col: 1, offset: 5998},
			expr: &actionExpr{
				pos: position{line: 267, col: 18, offset: 6015},
				run: (*parser).callonPrimitiveName1,
				expr: &labeledExpr{
					pos:   position{line: 267, col: 18, offset: 6015},
					label: "Name",
					expr: &ruleRefExpr{
						pos:  position{line: 267, col: 23, offset: 6020},
						name: "Identifier",
					},
				},
			},
		},
		{
			name: "Equation",
			pos:  position{line: 271, col: 1, offset: 6080},
			expr: &actionExpr{
				pos: position{line: 271, col: 13, offset: 6092},
				run: (*parser).callonEquation1,
				expr: &seqExpr{
					pos: position{line: 271, col: 13, offset: 6092},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 271, col: 13, offset: 6092},
							label: "First",
							expr: &ruleRefExpr{
								pos:  position{line: 271, col: 19, offset: 6098},
								name: "Constant",
							},
						},
						&seqExpr{
							pos: position{line: 271, col: 29, offset: 6108},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 271, col: 29, offset: 6108},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 271, col: 31, offset: 6110},
									val:        "^",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 271, col: 35, offset: 6114},
									name: "_",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 271, col: 38, offset: 6117},
							label: "Second",
							expr: &ruleRefExpr{
								pos:  position{line: 271, col: 45, offset: 6124},
								name: "Constant",
							},
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 283, col: 1, offset: 6273},
			expr: &choiceExpr{
				pos: position{line: 283, col: 10, offset: 6282},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 283, col: 10, offset: 6282},
						name: "Primitive",
					},
					&ruleRefExpr{
						pos:  position{line: 283, col: 20, offset: 6292},
						name: "Equation",
					},
					&ruleRefExpr{
						pos:  position{line: 283, col: 29, offset: 6301},
						name: "Constant",
					},
				},
			},
		},
		{
			name: "Queries",
			pos:  position{line: 285, col: 1, offset: 6312},
			expr: &actionExpr{
				pos: position{line: 285, col: 12, offset: 6323},
				run: (*parser).callonQueries1,
				expr: &seqExpr{
					pos: position{line: 285, col: 12, offset: 6323},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 285, col: 12, offset: 6323},
							val:        "queries",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 285, col: 22, offset: 6333},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 285, col: 24, offset: 6335},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 285, col: 28, offset: 6339},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 285, col: 30, offset: 6341},
							label: "Queries",
							expr: &zeroOrMoreExpr{
								pos: position{line: 285, col: 39, offset: 6350},
								expr: &ruleRefExpr{
									pos:  position{line: 285, col: 39, offset: 6350},
									name: "Query",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 285, col: 47, offset: 6358},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 285, col: 51, offset: 6362},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Query",
			pos:  position{line: 289, col: 1, offset: 6390},
			expr: &actionExpr{
				pos: position{line: 289, col: 10, offset: 6399},
				run: (*parser).callonQuery1,
				expr: &seqExpr{
					pos: position{line: 289, col: 10, offset: 6399},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 289, col: 10, offset: 6399},
							expr: &ruleRefExpr{
								pos:  position{line: 289, col: 10, offset: 6399},
								name: "Comment",
							},
						},
						&labeledExpr{
							pos:   position{line: 289, col: 19, offset: 6408},
							label: "Query",
							expr: &choiceExpr{
								pos: position{line: 289, col: 26, offset: 6415},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 289, col: 26, offset: 6415},
										name: "QueryConfidentiality",
									},
									&ruleRefExpr{
										pos:  position{line: 289, col: 47, offset: 6436},
										name: "QueryAuthentication",
									},
									&ruleRefExpr{
										pos:  position{line: 289, col: 67, offset: 6456},
										name: "QueryFreshness",
									},
									&ruleRefExpr{
										pos:  position{line: 289, col: 82, offset: 6471},
										name: "QueryUnlinkability",
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 289, col: 102, offset: 6491},
							expr: &ruleRefExpr{
								pos:  position{line: 289, col: 102, offset: 6491},
								name: "Comment",
							},
						},
					},
				},
			},
		},
		{
			name: "QueryConfidentiality",
			pos:  position{line: 293, col: 1, offset: 6525},
			expr: &actionExpr{
				pos: position{line: 293, col: 25, offset: 6549},
				run: (*parser).callonQueryConfidentiality1,
				expr: &seqExpr{
					pos: position{line: 293, col: 25, offset: 6549},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 293, col: 25, offset: 6549},
							val:        "confidentiality?",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 293, col: 44, offset: 6568},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 293, col: 46, offset: 6570},
							label: "Const",
							expr: &ruleRefExpr{
								pos:  position{line: 293, col: 52, offset: 6576},
								name: "Constant",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 293, col: 61, offset: 6585},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 293, col: 63, offset: 6587},
							label: "Options",
							expr: &zeroOrOneExpr{
								pos: position{line: 293, col: 71, offset: 6595},
								expr: &ruleRefExpr{
									pos:  position{line: 293, col: 71, offset: 6595},
									name: "QueryOptions",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 293, col: 85, offset: 6609},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "QueryAuthentication",
			pos:  position{line: 305, col: 1, offset: 6825},
			expr: &actionExpr{
				pos: position{line: 305, col: 24, offset: 6848},
				run: (*parser).callonQueryAuthentication1,
				expr: &seqExpr{
					pos: position{line: 305, col: 24, offset: 6848},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 305, col: 24, offset: 6848},
							val:        "authentication?",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 305, col: 42, offset: 6866},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 305, col: 44, offset: 6868},
							label: "Message",
							expr: &ruleRefExpr{
								pos:  position{line: 305, col: 52, offset: 6876},
								name: "Message",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 305, col: 60, offset: 6884},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 305, col: 62, offset: 6886},
							label: "Options",
							expr: &zeroOrOneExpr{
								pos: position{line: 305, col: 70, offset: 6894},
								expr: &ruleRefExpr{
									pos:  position{line: 305, col: 70, offset: 6894},
									name: "QueryOptions",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 305, col: 84, offset: 6908},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "QueryFreshness",
			pos:  position{line: 317, col: 1, offset: 7117},
			expr: &actionExpr{
				pos: position{line: 317, col: 19, offset: 7135},
				run: (*parser).callonQueryFreshness1,
				expr: &seqExpr{
					pos: position{line: 317, col: 19, offset: 7135},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 317, col: 19, offset: 7135},
							val:        "freshness?",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 317, col: 32, offset: 7148},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 317, col: 34, offset: 7150},
							label: "Const",
							expr: &ruleRefExpr{
								pos:  position{line: 317, col: 40, offset: 7156},
								name: "Constant",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 317, col: 49, offset: 7165},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 317, col: 51, offset: 7167},
							label: "Options",
							expr: &zeroOrOneExpr{
								pos: position{line: 317, col: 59, offset: 7175},
								expr: &ruleRefExpr{
									pos:  position{line: 317, col: 59, offset: 7175},
									name: "QueryOptions",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 317, col: 73, offset: 7189},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "QueryUnlinkability",
			pos:  position{line: 329, col: 1, offset: 7399},
			expr: &actionExpr{
				pos: position{line: 329, col: 23, offset: 7421},
				run: (*parser).callonQueryUnlinkability1,
				expr: &seqExpr{
					pos: position{line: 329, col: 23, offset: 7421},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 329, col: 23, offset: 7421},
							val:        "unlinkability?",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 329, col: 40, offset: 7438},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 329, col: 42, offset: 7440},
							label: "Constants",
							expr: &ruleRefExpr{
								pos:  position{line: 329, col: 52, offset: 7450},
								name: "Constants",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 329, col: 62, offset: 7460},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 329, col: 64, offset: 7462},
							label: "Options",
							expr: &zeroOrOneExpr{
								pos: position{line: 329, col: 72, offset: 7470},
								expr: &ruleRefExpr{
									pos:  position{line: 329, col: 72, offset: 7470},
									name: "QueryOptions",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 329, col: 86, offset: 7484},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "QueryOptions",
			pos:  position{line: 341, col: 1, offset: 7686},
			expr: &actionExpr{
				pos: position{line: 341, col: 17, offset: 7702},
				run: (*parser).callonQueryOptions1,
				expr: &seqExpr{
					pos: position{line: 341, col: 17, offset: 7702},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 341, col: 17, offset: 7702},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 341, col: 21, offset: 7706},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 341, col: 23, offset: 7708},
							label: "Options",
							expr: &zeroOrMoreExpr{
								pos: position{line: 341, col: 32, offset: 7717},
								expr: &ruleRefExpr{
									pos:  position{line: 341, col: 32, offset: 7717},
									name: "QueryOption",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 341, col: 46, offset: 7731},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 341, col: 50, offset: 7735},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "QueryOption",
			pos:  position{line: 348, col: 1, offset: 7872},
			expr: &actionExpr{
				pos: position{line: 348, col: 16, offset: 7887},
				run: (*parser).callonQueryOption1,
				expr: &seqExpr{
					pos: position{line: 348, col: 16, offset: 7887},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 348, col: 16, offset: 7887},
							label: "OptionName",
							expr: &ruleRefExpr{
								pos:  position{line: 348, col: 27, offset: 7898},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 348, col: 38, offset: 7909},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 348, col: 40, offset: 7911},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 348, col: 44, offset: 7915},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 348, col: 46, offset: 7917},
							label: "Message",
							expr: &ruleRefExpr{
								pos:  position{line: 348, col: 54, offset: 7925},
								name: "Message",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 348, col: 62, offset: 7933},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 348, col: 64, offset: 7935},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 348, col: 68, offset: 7939},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 355, col: 1, offset: 8042},
			expr: &actionExpr{
				pos: position{line: 355, col: 15, offset: 8056},
				run: (*parser).callonIdentifier1,
				expr: &labeledExpr{
					pos:   position{line: 355, col: 15, offset: 8056},
					label: "Identifier",
					expr: &oneOrMoreExpr{
						pos: position{line: 355, col: 26, offset: 8067},
						expr: &charClassMatcher{
							pos:        position{line: 355, col: 26, offset: 8067},
							val:        "[a-zA-Z0-9_]",
							chars:      []rune{'_'},
							ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 360, col: 1, offset: 8157},
			expr: &seqExpr{
				pos: position{line: 360, col: 12, offset: 8168},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 360, col: 12, offset: 8168},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 360, col: 14, offset: 8170},
						val:        "//",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 360, col: 19, offset: 8175},
						expr: &charClassMatcher{
							pos:        position{line: 360, col: 19, offset: 8175},
							val:        "[^\\n]",
							chars:      []rune{'\n'},
							ignoreCase: false,
							inverted:   true,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 360, col: 26, offset: 8182},
						name: "_",
					},
				},
			},
		},
		{
			name:        "_",
			displayName: "\"whitespace\"",
			pos:         position{line: 362, col: 1, offset: 8185},
			expr: &zeroOrMoreExpr{
				pos: position{line: 362, col: 19, offset: 8203},
				expr: &charClassMatcher{
					pos:        position{line: 362, col: 19, offset: 8203},
					val:        "[ \\t\\n\\r]",
					chars:      []rune{' ', '\t', '\n', '\r'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 364, col: 1, offset: 8215},
			expr: &notExpr{
				pos: position{line: 364, col: 8, offset: 8222},
				expr: &anyMatcher{
					line: 364, col: 9, offset: 8223,
				},
			},
		},
	},
}

func (c *current) onModel1(Attacker, Blocks, Queries interface{}) (interface{}, error) {
	switch {
	case Attacker == nil:
		return nil, errors.New("no `attacker` block defined")
	case Blocks == nil:
		return nil, errors.New("no principal or message blocks defined")
	case Queries == nil:
		return nil, errors.New("no `queries` block defined")
	}
	b := Blocks.([]interface{})
	q := Queries.([]interface{})
	db := make([]Block, len(b))
	dq := make([]Query, len(q))
	for i, v := range b {
		db[i] = v.(Block)
	}
	for i, v := range q {
		dq[i] = v.(Query)
	}
	return Model{
		Attacker: Attacker.(string),
		Blocks:   db,
		Queries:  dq,
	}, nil
}

func (p *parser) callonModel1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onModel1(stack["Attacker"], stack["Blocks"], stack["Queries"])
}

func (c *current) onAttacker1(Type interface{}) (interface{}, error) {
	return Type, nil
}

func (p *parser) callonAttacker1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAttacker1(stack["Type"])
}

func (c *current) onAttackerType1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonAttackerType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAttackerType1()
}

func (c *current) onBlock1(Block interface{}) (interface{}, error) {
	return Block, nil
}

func (p *parser) callonBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlock1(stack["Block"])
}

func (c *current) onPrincipal1(Name, Expressions interface{}) (interface{}, error) {
	e := Expressions.([]interface{})
	de := make([]Expression, len(e))
	for i, v := range e {
		de[i] = v.(Expression)
	}
	return Block{
		Kind: "principal",
		Principal: Principal{
			Name:        Name.(string),
			Expressions: de,
		},
	}, nil
}

func (p *parser) callonPrincipal1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPrincipal1(stack["Name"], stack["Expressions"])
}

func (c *current) onPrincipalName1(Name interface{}) (interface{}, error) {
	err := libpegCheckIfReserved(Name.(string))
	return strings.Title(Name.(string)), err
}

func (p *parser) callonPrincipalName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPrincipalName1(stack["Name"])
}

func (c *current) onQualifier1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonQualifier1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQualifier1()
}

func (c *current) onMessage1(Sender, Recipient, Constants interface{}) (interface{}, error) {
	return Block{
		Kind: "message",
		Message: Message{
			Sender:    Sender.(string),
			Recipient: Recipient.(string),
			Constants: Constants.([]Constant),
		},
	}, nil
}

func (p *parser) callonMessage1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMessage1(stack["Sender"], stack["Recipient"], stack["Constants"])
}

func (c *current) onMessageConstants1(MessageConstants interface{}) (interface{}, error) {
	var da []Constant
	a := MessageConstants.([]interface{})
	for _, v := range a {
		c := v.(Value).Constant
		da = append(da, c)
	}
	return da, nil
}

func (p *parser) callonMessageConstants1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMessageConstants1(stack["MessageConstants"])
}

func (c *current) onExpression1(Expression interface{}) (interface{}, error) {
	return Expression, nil
}

func (p *parser) callonExpression1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExpression1(stack["Expression"])
}

func (c *current) onKnows1(Qualifier, Constants interface{}) (interface{}, error) {
	return Expression{
		Kind:      "knows",
		Qualifier: Qualifier.(string),
		Constants: Constants.([]Constant),
	}, nil
}

func (p *parser) callonKnows1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onKnows1(stack["Qualifier"], stack["Constants"])
}

func (c *current) onGenerates1(Constants interface{}) (interface{}, error) {
	return Expression{
		Kind:      "generates",
		Qualifier: "",
		Constants: Constants.([]Constant),
	}, nil
}

func (p *parser) callonGenerates1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onGenerates1(stack["Constants"])
}

func (c *current) onLeaks1(Constants interface{}) (interface{}, error) {
	return Expression{
		Kind:      "leaks",
		Qualifier: "",
		Constants: Constants.([]Constant),
	}, nil
}

func (p *parser) callonLeaks1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLeaks1(stack["Constants"])
}

func (c *current) onAssignment1(Left, Right interface{}) (interface{}, error) {
	switch Right.(Value).Kind {
	case "constant":
		err := errors.New("cannot assign value to value")
		return nil, err
	}
	consts := Left.([]Constant)
	for i, c := range consts {
		switch c.Name {
		case "_":
			consts[i].Name = fmt.Sprintf("unnamed_%d", libpegUnnamedCounter)
			libpegUnnamedCounter = libpegUnnamedCounter + 1
		}
	}
	return Expression{
		Kind:      "assignment",
		Constants: consts,
		Assigned:  Right.(Value),
	}, nil
}

func (p *parser) callonAssignment1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignment1(stack["Left"], stack["Right"])
}

func (c *current) onConstant1(Const interface{}) (interface{}, error) {
	return Value{
		Kind: "constant",
		Constant: Constant{
			Name: Const.(string),
		},
	}, nil
}

func (p *parser) callonConstant1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConstant1(stack["Const"])
}

func (c *current) onConstants1(Constants interface{}) (interface{}, error) {
	var da []Constant
	var err error
	a := Constants.([]interface{})
	for _, c := range a {
		err = libpegCheckIfReserved(c.(Value).Constant.Name)
		if err != nil {
			break
		}
		da = append(da, c.(Value).Constant)
	}
	return da, err
}

func (p *parser) callonConstants1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConstants1(stack["Constants"])
}

func (c *current) onPhase1(Number interface{}) (interface{}, error) {
	a := Number.([]interface{})
	da := make([]uint8, len(a))
	for i, v := range a {
		da[i] = v.([]uint8)[0]
	}
	n, err := strconv.Atoi(b2s(da))
	return Block{
		Kind: "phase",
		Phase: Phase{
			Number: n,
		},
	}, err
}

func (p *parser) callonPhase1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPhase1(stack["Number"])
}

func (c *current) onGuardedConstant1(Guarded interface{}) (interface{}, error) {
	err := libpegCheckIfReserved(Guarded.(string))
	return Value{
		Kind: "constant",
		Constant: Constant{
			Name:  Guarded.(string),
			Guard: true,
		},
	}, err
}

func (p *parser) callonGuardedConstant1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onGuardedConstant1(stack["Guarded"])
}

func (c *current) onPrimitive1(Name, Arguments, Check interface{}) (interface{}, error) {
	args := []Value{}
	for _, a := range Arguments.([]interface{}) {
		args = append(args, a.(Value))
	}
	return Value{
		Kind: "primitive",
		Primitive: Primitive{
			Name:      Name.(string),
			Arguments: args,
			Output:    0,
			Check:     Check != nil,
		},
	}, nil
}

func (p *parser) callonPrimitive1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPrimitive1(stack["Name"], stack["Arguments"], stack["Check"])
}

func (c *current) onPrimitiveName1(Name interface{}) (interface{}, error) {
	return strings.ToUpper(Name.(string)), nil
}

func (p *parser) callonPrimitiveName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPrimitiveName1(stack["Name"])
}

func (c *current) onEquation1(First, Second interface{}) (interface{}, error) {
	return Value{
		Kind: "equation",
		Equation: Equation{
			Values: []Value{
				First.(Value),
				Second.(Value),
			},
		},
	}, nil
}

func (p *parser) callonEquation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEquation1(stack["First"], stack["Second"])
}

func (c *current) onQueries1(Queries interface{}) (interface{}, error) {
	return Queries, nil
}

func (p *parser) callonQueries1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQueries1(stack["Queries"])
}

func (c *current) onQuery1(Query interface{}) (interface{}, error) {
	return Query, nil
}

func (p *parser) callonQuery1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQuery1(stack["Query"])
}

func (c *current) onQueryConfidentiality1(Const, Options interface{}) (interface{}, error) {
	if Options == nil {
		Options = []QueryOption{}
	}
	return Query{
		Kind:      "confidentiality",
		Constants: []Constant{Const.(Value).Constant},
		Message:   Message{},
		Options:   Options.([]QueryOption),
	}, nil
}

func (p *parser) callonQueryConfidentiality1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQueryConfidentiality1(stack["Const"], stack["Options"])
}

func (c *current) onQueryAuthentication1(Message, Options interface{}) (interface{}, error) {
	if Options == nil {
		Options = []QueryOption{}
	}
	return Query{
		Kind:      "authentication",
		Constants: []Constant{},
		Message:   (Message.(Block)).Message,
		Options:   Options.([]QueryOption),
	}, nil
}

func (p *parser) callonQueryAuthentication1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQueryAuthentication1(stack["Message"], stack["Options"])
}

func (c *current) onQueryFreshness1(Const, Options interface{}) (interface{}, error) {
	if Options == nil {
		Options = []QueryOption{}
	}
	return Query{
		Kind:      "freshness",
		Constants: []Constant{Const.(Value).Constant},
		Message:   Message{},
		Options:   Options.([]QueryOption),
	}, nil
}

func (p *parser) callonQueryFreshness1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQueryFreshness1(stack["Const"], stack["Options"])
}

func (c *current) onQueryUnlinkability1(Constants, Options interface{}) (interface{}, error) {
	if Options == nil {
		Options = []QueryOption{}
	}
	return Query{
		Kind:      "unlinkability",
		Constants: Constants.([]Constant),
		Message:   Message{},
		Options:   Options.([]QueryOption),
	}, nil
}

func (p *parser) callonQueryUnlinkability1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQueryUnlinkability1(stack["Constants"], stack["Options"])
}

func (c *current) onQueryOptions1(Options interface{}) (interface{}, error) {
	o := Options.([]interface{})
	do := make([]QueryOption, len(o))
	for i, v := range o {
		do[i] = v.(QueryOption)
	}
	return do, nil
}

func (p *parser) callonQueryOptions1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQueryOptions1(stack["Options"])
}

func (c *current) onQueryOption1(OptionName, Message interface{}) (interface{}, error) {
	return QueryOption{
		Kind:    OptionName.(string),
		Message: (Message.(Block)).Message,
	}, nil
}

func (p *parser) callonQueryOption1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQueryOption1(stack["OptionName"], stack["Message"])
}

func (c *current) onIdentifier1(Identifier interface{}) (interface{}, error) {
	identifier := strings.ToLower(string(c.text))
	return identifier, nil
}

func (p *parser) callonIdentifier1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIdentifier1(stack["Identifier"])
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errNoMatch is returned if no match could be found.
	errNoMatch = errors.New("no match found")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner  error
	pos    position
	prefix string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	recover bool
	debug   bool
	depth   int

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position)
}

func (p *parser) addErrAt(err error, pos position) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String()}
	p.errs.add(pe)
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n == 1 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint
	var ok bool

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position)
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	// can't match EOF
	if cur == utf8.RuneError {
		return nil, false
	}
	start := p.pt
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(not.expr)
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}

func rangeTable(class string) *unicode.RangeTable {
	if rt, ok := unicode.Categories[class]; ok {
		return rt
	}
	if rt, ok := unicode.Properties[class]; ok {
		return rt
	}
	if rt, ok := unicode.Scripts[class]; ok {
		return rt
	}

	// cannot happen
	panic(fmt.Sprintf("invalid Unicode class: %s", class))
}
