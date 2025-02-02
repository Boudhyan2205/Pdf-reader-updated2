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

package sampling ;import (_cg "github.com/unidoc/unipdf/v3/internal/bitwise";_b "github.com/unidoc/unipdf/v3/internal/imageutil";_c "io";);func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _bd []uint32 ;_e :=bitsPerSample ;var _af uint32 ;
var _fc byte ;_cgc :=0;_ee :=0;_ge :=0;for _ge < len (data ){if _cgc > 0{_cbgd :=_cgc ;if _e < _cbgd {_cbgd =_e ;};_af =(_af <<uint (_cbgd ))|uint32 (_fc >>uint (8-_cbgd ));_cgc -=_cbgd ;if _cgc > 0{_fc =_fc <<uint (_cbgd );}else {_fc =0;};_e -=_cbgd ;
if _e ==0{_bd =append (_bd ,_af );_e =bitsPerSample ;_af =0;_ee ++;};}else {_eg :=data [_ge ];_ge ++;_bba :=8;if _e < _bba {_bba =_e ;};_cgc =8-_bba ;_af =(_af <<uint (_bba ))|uint32 (_eg >>uint (_cgc ));if _bba < 8{_fc =_eg <<uint (_bba );};_e -=_bba ;
if _e ==0{_bd =append (_bd ,_af );_e =bitsPerSample ;_af =0;_ee ++;};};};for _cgc >=bitsPerSample {_ff :=_cgc ;if _e < _ff {_ff =_e ;};_af =(_af <<uint (_ff ))|uint32 (_fc >>uint (8-_ff ));_cgc -=_ff ;if _cgc > 0{_fc =_fc <<uint (_ff );}else {_fc =0;};
_e -=_ff ;if _e ==0{_bd =append (_bd ,_af );_e =bitsPerSample ;_af =0;_ee ++;};};return _bd ;};func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _gca []uint32 ;_ab :=bitsPerOutputSample ;var _dg uint32 ;
var _cge uint32 ;_cc :=0;_ag :=0;_fcb :=0;for _fcb < len (data ){if _cc > 0{_fg :=_cc ;if _ab < _fg {_fg =_ab ;};_dg =(_dg <<uint (_fg ))|(_cge >>uint (bitsPerInputSample -_fg ));_cc -=_fg ;if _cc > 0{_cge =_cge <<uint (_fg );}else {_cge =0;};_ab -=_fg ;
if _ab ==0{_gca =append (_gca ,_dg );_ab =bitsPerOutputSample ;_dg =0;_ag ++;};}else {_ga :=data [_fcb ];_fcb ++;_ba :=bitsPerInputSample ;if _ab < _ba {_ba =_ab ;};_cc =bitsPerInputSample -_ba ;_dg =(_dg <<uint (_ba ))|(_ga >>uint (_cc ));if _ba < bitsPerInputSample {_cge =_ga <<uint (_ba );
};_ab -=_ba ;if _ab ==0{_gca =append (_gca ,_dg );_ab =bitsPerOutputSample ;_dg =0;_ag ++;};};};for _cc >=bitsPerOutputSample {_afa :=_cc ;if _ab < _afa {_afa =_ab ;};_dg =(_dg <<uint (_afa ))|(_cge >>uint (bitsPerInputSample -_afa ));_cc -=_afa ;if _cc > 0{_cge =_cge <<uint (_afa );
}else {_cge =0;};_ab -=_afa ;if _ab ==0{_gca =append (_gca ,_dg );_ab =bitsPerOutputSample ;_dg =0;_ag ++;};};if _ab > 0&&_ab < bitsPerOutputSample {_dg <<=uint (_ab );_gca =append (_gca ,_dg );};return _gca ;};type SampleReader interface{ReadSample ()(uint32 ,error );
ReadSamples (_bb []uint32 )error ;};func NewReader (img _b .ImageBase )*Reader {return &Reader {_dc :_cg .NewReader (img .Data ),_d :img ,_cb :img .ColorComponents ,_fd :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};
func NewWriter (img _b .ImageBase )*Writer {return &Writer {_dcf :_cg .NewWriterMSB (img .Data ),_gb :img ,_gd :img .ColorComponents ,_dga :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};func (_fda *Writer )WriteSample (sample uint32 )error {if _ ,_ef :=_fda ._dcf .WriteBits (uint64 (sample ),_fda ._gb .BitsPerComponent );
_ef !=nil {return _ef ;};_fda ._gd --;if _fda ._gd ==0{_fda ._gd =_fda ._gb .ColorComponents ;_fda ._ac ++;};if _fda ._ac ==_fda ._gb .Width {if _fda ._dga {_fda ._dcf .FinishByte ();};_fda ._ac =0;};return nil ;};type Reader struct{_d _b .ImageBase ;_dc *_cg .Reader ;
_g ,_a ,_cb int ;_fd bool ;};type Writer struct{_gb _b .ImageBase ;_dcf *_cg .Writer ;_ac ,_gd int ;_dga bool ;};func (_cga *Writer )WriteSamples (samples []uint32 )error {for _gdg :=0;_gdg < len (samples );_gdg ++{if _bdg :=_cga .WriteSample (samples [_gdg ]);
_bdg !=nil {return _bdg ;};};return nil ;};func (_bg *Reader )ReadSample ()(uint32 ,error ){if _bg ._a ==_bg ._d .Height {return 0,_c .EOF ;};_ca ,_fe :=_bg ._dc .ReadBits (byte (_bg ._d .BitsPerComponent ));if _fe !=nil {return 0,_fe ;};_bg ._cb --;if _bg ._cb ==0{_bg ._cb =_bg ._d .ColorComponents ;
_bg ._g ++;};if _bg ._g ==_bg ._d .Width {if _bg ._fd {_bg ._dc .ConsumeRemainingBits ();};_bg ._g =0;_bg ._a ++;};return uint32 (_ca ),nil ;};func (_gc *Reader )ReadSamples (samples []uint32 )(_cag error ){for _cbg :=0;_cbg < len (samples );_cbg ++{samples [_cbg ],_cag =_gc .ReadSample ();
if _cag !=nil {return _cag ;};};return nil ;};type SampleWriter interface{WriteSample (_fa uint32 )error ;WriteSamples (_bab []uint32 )error ;};