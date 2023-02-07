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

package textshaping ;import (_e "github.com/unidoc/garabic";_eg "golang.org/x/text/unicode/bidi";_c "strings";);

// ArabicShape returns shaped arabic glyphs string.
func ArabicShape (text string )(string ,error ){_ac :=_eg .Paragraph {};_ac .SetString (text );_f ,_aa :=_ac .Order ();if _aa !=nil {return "",_aa ;};for _d :=0;_d < _f .NumRuns ();_d ++{_db :=_f .Run (_d );_g :=_db .String ();if _db .Direction ()==_eg .RightToLeft {var (_b =_e .Shape (_g );
_fb =[]rune (_b );_dc =make ([]rune ,len (_fb )););_ege :=0;for _ea :=len (_fb )-1;_ea >=0;_ea --{_dc [_ege ]=_fb [_ea ];_ege ++;};_g =string (_dc );text =_c .Replace (text ,_c .TrimSpace (_db .String ()),_g ,1);};};return text ,nil ;};