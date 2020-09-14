//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package errors ;import _d "fmt";func _gf (_b ,_ec string )*processError {return &processError {_ee :"\u005b\u0055\u006e\u0069\u0050\u0044\u0046\u005d",_a :_b ,_dd :_ec };};func Wrap (err error ,processName ,message string )error {if _ecb ,_ed :=err .(*processError );_ed {_ecb ._ee ="";};_ef :=_gf (message ,processName );_ef ._ab =err ;return _ef ;};func Wrapf (err error ,processName ,message string ,arguments ...interface{})error {if _eca ,_efg :=err .(*processError );_efg {_eca ._ee ="";};_de :=_gf (_d .Sprintf (message ,arguments ...),processName );_de ._ab =err ;return _de ;};func (_g *processError )Error ()string {var _f string ;if _g ._ee !=""{_f =_g ._ee ;};_f +="\u0050r\u006f\u0063\u0065\u0073\u0073\u003a "+_g ._dd ;if _g ._a !=""{_f +="\u0020\u004d\u0065\u0073\u0073\u0061\u0067\u0065\u003a\u0020"+_g ._a ;};if _g ._ab !=nil {_f +="\u002e\u0020"+_g ._ab .Error ();};return _f ;};func Error (processName ,message string )error {return _gf (message ,processName )};func Errorf (processName ,message string ,arguments ...interface{})error {return _gf (_d .Sprintf (message ,arguments ...),processName );};type processError struct{_ee string ;_dd string ;_a string ;_ab error ;};