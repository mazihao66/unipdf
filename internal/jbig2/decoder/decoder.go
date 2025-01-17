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

package decoder ;import (_e "github.com/unidoc/unipdf/v3/internal/bitwise";_g "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_ce "github.com/unidoc/unipdf/v3/internal/jbig2/document";_gb "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_f "image";
);func (_ge *Decoder )decodePageImage (_ag int )(_f .Image ,error ){const _ged ="\u0064e\u0063o\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";if _ag < 0{return nil ,_gb .Errorf (_ged ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_ag );
};if _ag > int (_ge ._fg .NumberOfPages ){return nil ,_gb .Errorf (_ged ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_ag );
};_fea ,_ba :=_ge ._fg .GetPage (_ag );if _ba !=nil {return nil ,_gb .Wrap (_ba ,_ged ,"");};_feb ,_ba :=_fea .GetBitmap ();if _ba !=nil {return nil ,_gb .Wrap (_ba ,_ged ,"");};_feb .InverseData ();return _feb .ToImage (),nil ;};func Decode (input []byte ,parameters Parameters ,globals *_ce .Globals )(*Decoder ,error ){_bf :=_e .NewReader (input );
_ff ,_gf :=_ce .DecodeDocument (_bf ,globals );if _gf !=nil {return nil ,_gf ;};return &Decoder {_b :_bf ,_fg :_ff ,_cc :parameters },nil ;};func (_a *Decoder )DecodePage (pageNumber int )([]byte ,error ){return _a .decodePage (pageNumber )};func (_gd *Decoder )decodePage (_ee int )([]byte ,error ){const _bc ="\u0064\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065";
if _ee < 0{return nil ,_gb .Errorf (_bc ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_ee );};if _ee > int (_gd ._fg .NumberOfPages ){return nil ,_gb .Errorf (_bc ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_ee );
};_fa ,_gbf :=_gd ._fg .GetPage (_ee );if _gbf !=nil {return nil ,_gb .Wrap (_gbf ,_bc ,"");};_bbd ,_gbf :=_fa .GetBitmap ();if _gbf !=nil {return nil ,_gb .Wrap (_gbf ,_bc ,"");};_bbd .InverseData ();if !_gd ._cc .UnpaddedData {return _bbd .Data ,nil ;
};return _bbd .GetUnpaddedData ();};type Parameters struct{UnpaddedData bool ;Color _g .Color ;};func (_ca *Decoder )PageNumber ()(int ,error ){const _eg ="\u0044e\u0063o\u0064\u0065\u0072\u002e\u0050a\u0067\u0065N\u0075\u006d\u0062\u0065\u0072";if _ca ._fg ==nil {return 0,_gb .Error (_eg ,"d\u0065\u0063\u006f\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0069\u006e\u0069\u0074\u0069a\u006c\u0069\u007ae\u0064 \u0079\u0065\u0074");
};return int (_ca ._fg .NumberOfPages ),nil ;};func (_eb *Decoder )DecodeNextPage ()([]byte ,error ){_eb ._d ++;_fe :=_eb ._d ;return _eb .decodePage (_fe );};type Decoder struct{_b _e .StreamReader ;_fg *_ce .Document ;_d int ;_cc Parameters ;};func (_fgf *Decoder )DecodePageImage (pageNumber int )(_f .Image ,error ){const _bb ="\u0064\u0065\u0063od\u0065\u0072\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";
_de ,_ec :=_fgf .decodePageImage (pageNumber );if _ec !=nil {return nil ,_gb .Wrap (_ec ,_bb ,"");};return _de ,nil ;};