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

package sampling ;import (_g "github.com/unidoc/unipdf/v3/internal/bitwise";_d "github.com/unidoc/unipdf/v3/internal/imageutil";_e "io";);func NewWriter (img _d .ImageBase )*Writer {return &Writer {_ga :_g .NewWriterMSB (img .Data ),_debf :img ,_bff :img .ColorComponents ,_gae :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };
};type Reader struct{_bf _d .ImageBase ;_gc *_g .Reader ;_f ,_a ,_da int ;_ea bool ;};func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _ebg []uint32 ;_ba :=bitsPerOutputSample ;var _fg uint32 ;var _cf uint32 ;
_gf :=0;_dg :=0;_cfb :=0;for _cfb < len (data ){if _gf > 0{_cfe :=_gf ;if _ba < _cfe {_cfe =_ba ;};_fg =(_fg <<uint (_cfe ))|(_cf >>uint (bitsPerInputSample -_cfe ));_gf -=_cfe ;if _gf > 0{_cf =_cf <<uint (_cfe );}else {_cf =0;};_ba -=_cfe ;if _ba ==0{_ebg =append (_ebg ,_fg );
_ba =bitsPerOutputSample ;_fg =0;_dg ++;};}else {_ag :=data [_cfb ];_cfb ++;_eba :=bitsPerInputSample ;if _ba < _eba {_eba =_ba ;};_gf =bitsPerInputSample -_eba ;_fg =(_fg <<uint (_eba ))|(_ag >>uint (_gf ));if _eba < bitsPerInputSample {_cf =_ag <<uint (_eba );
};_ba -=_eba ;if _ba ==0{_ebg =append (_ebg ,_fg );_ba =bitsPerOutputSample ;_fg =0;_dg ++;};};};for _gf >=bitsPerOutputSample {_gff :=_gf ;if _ba < _gff {_gff =_ba ;};_fg =(_fg <<uint (_gff ))|(_cf >>uint (bitsPerInputSample -_gff ));_gf -=_gff ;if _gf > 0{_cf =_cf <<uint (_gff );
}else {_cf =0;};_ba -=_gff ;if _ba ==0{_ebg =append (_ebg ,_fg );_ba =bitsPerOutputSample ;_fg =0;_dg ++;};};if _ba > 0&&_ba < bitsPerOutputSample {_fg <<=uint (_ba );_ebg =append (_ebg ,_fg );};return _ebg ;};func NewReader (img _d .ImageBase )*Reader {return &Reader {_gc :_g .NewReader (img .Data ),_bf :img ,_da :img .ColorComponents ,_ea :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };
};func (_fa *Reader )ReadSample ()(uint32 ,error ){if _fa ._a ==_fa ._bf .Height {return 0,_e .EOF ;};_eab ,_bd :=_fa ._gc .ReadBits (byte (_fa ._bf .BitsPerComponent ));if _bd !=nil {return 0,_bd ;};_fa ._da --;if _fa ._da ==0{_fa ._da =_fa ._bf .ColorComponents ;
_fa ._f ++;};if _fa ._f ==_fa ._bf .Width {if _fa ._ea {_fa ._gc .ConsumeRemainingBits ();};_fa ._f =0;_fa ._a ++;};return uint32 (_eab ),nil ;};type SampleReader interface{ReadSample ()(uint32 ,error );ReadSamples (_ee []uint32 )error ;};type SampleWriter interface{WriteSample (_dge uint32 )error ;
WriteSamples (_gcd []uint32 )error ;};func (_eb *Reader )ReadSamples (samples []uint32 )(_ed error ){for _de :=0;_de < len (samples );_de ++{samples [_de ],_ed =_eb .ReadSample ();if _ed !=nil {return _ed ;};};return nil ;};func (_dc *Writer )WriteSample (sample uint32 )error {if _ ,_ge :=_dc ._ga .WriteBits (uint64 (sample ),_dc ._debf .BitsPerComponent );
_ge !=nil {return _ge ;};_dc ._bff --;if _dc ._bff ==0{_dc ._bff =_dc ._debf .ColorComponents ;_dc ._acb ++;};if _dc ._acb ==_dc ._debf .Width {if _dc ._gae {_dc ._ga .FinishByte ();};_dc ._acb =0;};return nil ;};func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _ac []uint32 ;
_gg :=bitsPerSample ;var _fe uint32 ;var _ec byte ;_fec :=0;_c :=0;_gb :=0;for _gb < len (data ){if _fec > 0{_fecc :=_fec ;if _gg < _fecc {_fecc =_gg ;};_fe =(_fe <<uint (_fecc ))|uint32 (_ec >>uint (8-_fecc ));_fec -=_fecc ;if _fec > 0{_ec =_ec <<uint (_fecc );
}else {_ec =0;};_gg -=_fecc ;if _gg ==0{_ac =append (_ac ,_fe );_gg =bitsPerSample ;_fe =0;_c ++;};}else {_ad :=data [_gb ];_gb ++;_deb :=8;if _gg < _deb {_deb =_gg ;};_fec =8-_deb ;_fe =(_fe <<uint (_deb ))|uint32 (_ad >>uint (_fec ));if _deb < 8{_ec =_ad <<uint (_deb );
};_gg -=_deb ;if _gg ==0{_ac =append (_ac ,_fe );_gg =bitsPerSample ;_fe =0;_c ++;};};};for _fec >=bitsPerSample {_ef :=_fec ;if _gg < _ef {_ef =_gg ;};_fe =(_fe <<uint (_ef ))|uint32 (_ec >>uint (8-_ef ));_fec -=_ef ;if _fec > 0{_ec =_ec <<uint (_ef );
}else {_ec =0;};_gg -=_ef ;if _gg ==0{_ac =append (_ac ,_fe );_gg =bitsPerSample ;_fe =0;_c ++;};};return _ac ;};func (_ce *Writer )WriteSamples (samples []uint32 )error {for _gfg :=0;_gfg < len (samples );_gfg ++{if _ae :=_ce .WriteSample (samples [_gfg ]);
_ae !=nil {return _ae ;};};return nil ;};type Writer struct{_debf _d .ImageBase ;_ga *_g .Writer ;_acb ,_bff int ;_gae bool ;};