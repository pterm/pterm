// Package putils converts common Go values into the data pterm printers
// take, and wraps everyday jobs such as running a function behind a
// spinner.
//
// Every converter returns plain pterm types, so the result works with any
// printer configuration:
//
//	pterm.Table.WithData(putils.TableFromStructs(users)).Render()
//	pterm.Tree.WithRoot(putils.TreeFromMap("config", cfg)).Render()
//	pterm.KeyValue.WithPairs(putils.KeyValueFromStruct(build)...).Render()
package putils
