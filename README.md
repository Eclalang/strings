# Strings library

[![Go Report Card](https://goreportcard.com/badge/github.com/Eclalang/strings)](https://goreportcard.com/report/github.com/Eclalang/strings)
[![codecov](https://codecov.io/gh/Eclalang/strings/graph/badge.svg?token=YNCIYERVBO)](https://codecov.io/gh/Eclalang/strings)

## Candidate functions :

|  Func Name  |                               Prototype                                |                                                                                  Description                                                                                  | Comments |
|:-----------:|:----------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|:--------:|
|  Contains   |              Contains(str string, substr string) bool {}               |                                                               Returns true if the string contains the substring                                                               |    N/A   |
| ContainsAny |             ContainsAny(str string, chars string) bool {}              |                                                           Returns true if the string contains any of the characters                                                           |    N/A   |
|    Count    |                Count(str string, substr string) int {}                 |                                                       Returns the number of non-overlapping instances of substr in str                                                        |    N/A   |
|     Cut     | Cut(str string, sep string) before string, after string, found bool {} |                                               Returns a string before and after the separator, and a bool if it's found or not                                                |    N/A   |
|  HasPrefix  |              HasPrefix(str string, prefix string) bool {}              |                                                                Returns true if the string starts by the prefix                                                                |    N/A   |
|  HasSuffix  |              HasSuffix(str string, suffix string) bool {}              |                                                                 Returns true if the string ends by the suffix                                                                 |    N/A   |
|   IndexOf   |               IndexOf(str string, substr string) int {}                |                                                 Returns the index of the first instance of substr in str, or -1 if not found                                                  |    N/A   |
|    Join     |               Join(elems []string, sep string) string {}               |                                                    Returns a concatenated string from an array of string separated by sep                                                     |    N/A   |
|   Replace   |     Replace(str string, old string, new string, nb int) string {}      |                                                        Returns a string with the first instance of old replaced by new                                                        |    N/A   |
| ReplaceAll  |        ReplaceAll(str string, old string, new string) string {}        |                                                          Returns a string with all instances of old replaced by new                                                           |    N/A   |
|    Split    |               Split(str string, sep string) []string {}                |                             Returns an array of the substrings between the separator, or an array only containing str if it doesn't contains sep                              |    N/A   |
| SplitAfter  |             SplitAfter(str string, sep string) []string {}             |                              Returns an array of the substrings after the separator, or an array only containing str if it doesn't contains sep                               |    N/A   |
| SplitAfterN |        SplitAfterN(str string, sep string, nb int) []string {}         |  Returns an array of the substrings after the separator, or an array only containing str if it doesn't contains sep. The count determines the number of substrings to return  |    N/A   |
|   SplitN    |           SplitN(str string, sep string, nb int) []string {}           | Returns an array of the substrings between the separator, or an array only containing str if it doesn't contains sep. The count determines the number of substrings to return |    N/A   |
|   ToLower   |                     ToLower(str string) string {}                      |                                                               Returns a string with all characters in lowercase                                                               |    N/A   |
|   ToUpper   |                     ToUpper(str string) string {}                      |                                                               Returns a string with all characters in uppercase                                                               |    N/A   |
|    Trim     |               Trim(str string, cutset string) string {}                |                                                               Returns a string with all cut characters removed                                                                |    N/A   |
