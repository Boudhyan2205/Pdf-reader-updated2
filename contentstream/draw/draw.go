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

// Package draw has handy features for defining paths which can be used to draw content on a PDF page.  Handles
// defining paths as points, vector calculations and conversion to PDF content stream data which can be used in
// page content streams and XObject forms and thus also in annotation appearance streams.
//
// Also defines utility functions for drawing common shapes such as rectangles, lines and circles (ovals).
package draw ;import (_ee "fmt";_df "github.com/unidoc/unipdf/v3/contentstream";_f "github.com/unidoc/unipdf/v3/core";_eb "github.com/unidoc/unipdf/v3/internal/transform";_a "github.com/unidoc/unipdf/v3/model";_e "math";);

// LineEndingStyle defines the line ending style for lines.
// The currently supported line ending styles are None, Arrow (ClosedArrow) and Butt.
type LineEndingStyle int ;

// PolyBezierCurve represents a composite curve that is the result of
// joining multiple cubic Bezier curves.
type PolyBezierCurve struct{Curves []CubicBezierCurve ;BorderWidth float64 ;BorderColor *_a .PdfColorDeviceRGB ;FillEnabled bool ;FillColor *_a .PdfColorDeviceRGB ;};

// GetPolarAngle returns the angle the magnitude of the vector forms with the
// positive X-axis going counterclockwise.
func (_egab Vector )GetPolarAngle ()float64 {return _e .Atan2 (_egab .Dy ,_egab .Dx )};

// RemovePoint removes the point at the index specified by number from the
// path. The index is 1-based.
func (_fcb Path )RemovePoint (number int )Path {if number < 1||number > len (_fcb .Points ){return _fcb ;};_bcg :=number -1;_fcb .Points =append (_fcb .Points [:_bcg ],_fcb .Points [_bcg +1:]...);return _fcb ;};

// Draw draws the polygon. A graphics state name can be specified for
// setting the polygon properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polygon
// bounding box.
func (_ffe Polygon )Draw (gsName string )([]byte ,*_a .PdfRectangle ,error ){_ebc :=_df .NewContentCreator ();_ebc .Add_q ();_ffe .FillEnabled =_ffe .FillEnabled &&_ffe .FillColor !=nil ;if _ffe .FillEnabled {_ebc .Add_rg (_ffe .FillColor .R (),_ffe .FillColor .G (),_ffe .FillColor .B ());};_ffe .BorderEnabled =_ffe .BorderEnabled &&_ffe .BorderColor !=nil ;if _ffe .BorderEnabled {_ebc .Add_RG (_ffe .BorderColor .R (),_ffe .BorderColor .G (),_ffe .BorderColor .B ());_ebc .Add_w (_ffe .BorderWidth );};if len (gsName )> 1{_ebc .Add_gs (_f .PdfObjectName (gsName ));};_bfd :=NewPath ();for _ ,_fbd :=range _ffe .Points {for _bdc ,_eab :=range _fbd {_bfd =_bfd .AppendPoint (_eab );if _bdc ==0{_ebc .Add_m (_eab .X ,_eab .Y );}else {_ebc .Add_l (_eab .X ,_eab .Y );};};_ebc .Add_h ();};if _ffe .FillEnabled &&_ffe .BorderEnabled {_ebc .Add_B ();}else if _ffe .FillEnabled {_ebc .Add_f ();}else if _ffe .BorderEnabled {_ebc .Add_S ();};_ebc .Add_Q ();return _ebc .Bytes (),_bfd .GetBoundingBox ().ToPdfRectangle (),nil ;};

// Copy returns a clone of the Bezier path.
func (_cc CubicBezierPath )Copy ()CubicBezierPath {_ff :=CubicBezierPath {};_ff .Curves =[]CubicBezierCurve {};for _ ,_dc :=range _cc .Curves {_ff .Curves =append (_ff .Curves ,_dc );};return _ff ;};

// Point represents a two-dimensional point.
type Point struct{X float64 ;Y float64 ;};

// ToPdfRectangle returns the bounding box as a PDF rectangle.
func (_bac BoundingBox )ToPdfRectangle ()*_a .PdfRectangle {return &_a .PdfRectangle {Llx :_bac .X ,Lly :_bac .Y ,Urx :_bac .X +_bac .Width ,Ury :_bac .Y +_bac .Height };};

// ToPdfRectangle returns the rectangle as a PDF rectangle.
func (_bgc Rectangle )ToPdfRectangle ()*_a .PdfRectangle {return &_a .PdfRectangle {Llx :_bgc .X ,Lly :_bgc .Y ,Urx :_bgc .X +_bgc .Width ,Ury :_bgc .Y +_bgc .Height };};

// Rotate returns a new Point at `p` rotated by `theta` degrees.
func (_aa Point )Rotate (theta float64 )Point {_bgaa :=_eb .NewPoint (_aa .X ,_aa .Y ).Rotate (theta );return NewPoint (_bgaa .X ,_bgaa .Y );};func (_fcbb Point )String ()string {return _ee .Sprintf ("(\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u0029",_fcbb .X ,_fcbb .Y );};

// DrawPathWithCreator makes the path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawPathWithCreator (path Path ,creator *_df .ContentCreator ){for _def ,_ebd :=range path .Points {if _def ==0{creator .Add_m (_ebd .X ,_ebd .Y );}else {creator .Add_l (_ebd .X ,_ebd .Y );};};};

// NewCubicBezierPath returns a new empty cubic Bezier path.
func NewCubicBezierPath ()CubicBezierPath {_fe :=CubicBezierPath {};_fe .Curves =[]CubicBezierCurve {};return _fe ;};

// Polyline defines a slice of points that are connected as straight lines.
type Polyline struct{Points []Point ;LineColor *_a .PdfColorDeviceRGB ;LineWidth float64 ;};

// LineStyle refers to how the line will be created.
type LineStyle int ;

// GetBounds returns the bounding box of the Bezier curve.
func (_cg CubicBezierCurve )GetBounds ()_a .PdfRectangle {_db :=_cg .P0 .X ;_dg :=_cg .P0 .X ;_b :=_cg .P0 .Y ;_ad :=_cg .P0 .Y ;for _eg :=0.0;_eg <=1.0;_eg +=0.001{Rx :=_cg .P0 .X *_e .Pow (1-_eg ,3)+_cg .P1 .X *3*_eg *_e .Pow (1-_eg ,2)+_cg .P2 .X *3*_e .Pow (_eg ,2)*(1-_eg )+_cg .P3 .X *_e .Pow (_eg ,3);Ry :=_cg .P0 .Y *_e .Pow (1-_eg ,3)+_cg .P1 .Y *3*_eg *_e .Pow (1-_eg ,2)+_cg .P2 .Y *3*_e .Pow (_eg ,2)*(1-_eg )+_cg .P3 .Y *_e .Pow (_eg ,3);if Rx < _db {_db =Rx ;};if Rx > _dg {_dg =Rx ;};if Ry < _b {_b =Ry ;};if Ry > _ad {_ad =Ry ;};};_g :=_a .PdfRectangle {};_g .Llx =_db ;_g .Lly =_b ;_g .Urx =_dg ;_g .Ury =_ad ;return _g ;};

// NewCubicBezierCurve returns a new cubic Bezier curve.
func NewCubicBezierCurve (x0 ,y0 ,x1 ,y1 ,x2 ,y2 ,x3 ,y3 float64 )CubicBezierCurve {_c :=CubicBezierCurve {};_c .P0 =NewPoint (x0 ,y0 );_c .P1 =NewPoint (x1 ,y1 );_c .P2 =NewPoint (x2 ,y2 );_c .P3 =NewPoint (x3 ,y3 );return _c ;};

// Length returns the number of points in the path.
func (_gf Path )Length ()int {return len (_gf .Points )};

// AddVector adds vector to a point.
func (_bga Point )AddVector (v Vector )Point {_bga .X +=v .Dx ;_bga .Y +=v .Dy ;return _bga };

// Flip changes the sign of the vector: -vector.
func (_cb Vector )Flip ()Vector {_gdb :=_cb .Magnitude ();_bfb :=_cb .GetPolarAngle ();_cb .Dx =_gdb *_e .Cos (_bfb +_e .Pi );_cb .Dy =_gdb *_e .Sin (_bfb +_e .Pi );return _cb ;};

// Draw draws the circle. Can specify a graphics state (gsName) for setting opacity etc.  Otherwise leave empty ("").
// Returns the content stream as a byte array, the bounding box and an error on failure.
func (_ab Circle )Draw (gsName string )([]byte ,*_a .PdfRectangle ,error ){_cga :=_ab .Width /2;_ge :=_ab .Height /2;if _ab .BorderEnabled {_cga -=_ab .BorderWidth /2;_ge -=_ab .BorderWidth /2;};_dfg :=0.551784;_dbb :=_cga *_dfg ;_ggad :=_ge *_dfg ;_de :=NewCubicBezierPath ();_de =_de .AppendCurve (NewCubicBezierCurve (-_cga ,0,-_cga ,_ggad ,-_dbb ,_ge ,0,_ge ));_de =_de .AppendCurve (NewCubicBezierCurve (0,_ge ,_dbb ,_ge ,_cga ,_ggad ,_cga ,0));_de =_de .AppendCurve (NewCubicBezierCurve (_cga ,0,_cga ,-_ggad ,_dbb ,-_ge ,0,-_ge ));_de =_de .AppendCurve (NewCubicBezierCurve (0,-_ge ,-_dbb ,-_ge ,-_cga ,-_ggad ,-_cga ,0));_de =_de .Offset (_cga ,_ge );if _ab .BorderEnabled {_de =_de .Offset (_ab .BorderWidth /2,_ab .BorderWidth /2);};if _ab .X !=0||_ab .Y !=0{_de =_de .Offset (_ab .X ,_ab .Y );};_ead :=_df .NewContentCreator ();_ead .Add_q ();if _ab .FillEnabled {_ead .Add_rg (_ab .FillColor .R (),_ab .FillColor .G (),_ab .FillColor .B ());};if _ab .BorderEnabled {_ead .Add_RG (_ab .BorderColor .R (),_ab .BorderColor .G (),_ab .BorderColor .B ());_ead .Add_w (_ab .BorderWidth );};if len (gsName )> 1{_ead .Add_gs (_f .PdfObjectName (gsName ));};DrawBezierPathWithCreator (_de ,_ead );_ead .Add_h ();if _ab .FillEnabled &&_ab .BorderEnabled {_ead .Add_B ();}else if _ab .FillEnabled {_ead .Add_f ();}else if _ab .BorderEnabled {_ead .Add_S ();};_ead .Add_Q ();_dee :=_de .GetBoundingBox ();if _ab .BorderEnabled {_dee .Height +=_ab .BorderWidth ;_dee .Width +=_ab .BorderWidth ;_dee .X -=_ab .BorderWidth /2;_dee .Y -=_ab .BorderWidth /2;};return _ead .Bytes (),_dee .ToPdfRectangle (),nil ;};

// Draw draws the line to PDF contentstream. Generates the content stream which can be used in page contents or
// appearance stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error
// if one occurred.
func (_fgf Line )Draw (gsName string )([]byte ,*_a .PdfRectangle ,error ){_aff ,_dff :=_fgf .X1 ,_fgf .X2 ;_ged ,_dea :=_fgf .Y1 ,_fgf .Y2 ;_cge :=_dea -_ged ;_gc :=_dff -_aff ;_eba :=_e .Atan2 (_cge ,_gc );L :=_e .Sqrt (_e .Pow (_gc ,2.0)+_e .Pow (_cge ,2.0));_fa :=_fgf .LineWidth ;_dge :=_e .Pi ;_bace :=1.0;if _gc < 0{_bace *=-1.0;};if _cge < 0{_bace *=-1.0;};VsX :=_bace *(-_fa /2*_e .Cos (_eba +_dge /2));VsY :=_bace *(-_fa /2*_e .Sin (_eba +_dge /2)+_fa *_e .Sin (_eba +_dge /2));V1X :=VsX +_fa /2*_e .Cos (_eba +_dge /2);V1Y :=VsY +_fa /2*_e .Sin (_eba +_dge /2);V2X :=VsX +_fa /2*_e .Cos (_eba +_dge /2)+L *_e .Cos (_eba );V2Y :=VsY +_fa /2*_e .Sin (_eba +_dge /2)+L *_e .Sin (_eba );V3X :=VsX +_fa /2*_e .Cos (_eba +_dge /2)+L *_e .Cos (_eba )+_fa *_e .Cos (_eba -_dge /2);V3Y :=VsY +_fa /2*_e .Sin (_eba +_dge /2)+L *_e .Sin (_eba )+_fa *_e .Sin (_eba -_dge /2);V4X :=VsX +_fa /2*_e .Cos (_eba -_dge /2);V4Y :=VsY +_fa /2*_e .Sin (_eba -_dge /2);_dbad :=NewPath ();_dbad =_dbad .AppendPoint (NewPoint (V1X ,V1Y ));_dbad =_dbad .AppendPoint (NewPoint (V2X ,V2Y ));_dbad =_dbad .AppendPoint (NewPoint (V3X ,V3Y ));_dbad =_dbad .AppendPoint (NewPoint (V4X ,V4Y ));_eeb :=_fgf .LineEndingStyle1 ;_aba :=_fgf .LineEndingStyle2 ;_gb :=3*_fa ;_deee :=3*_fa ;_ggg :=(_deee -_fa )/2;if _aba ==LineEndingStyleArrow {_ca :=_dbad .GetPointNumber (2);_adbb :=NewVectorPolar (_gb ,_eba +_dge );_caf :=_ca .AddVector (_adbb );_gfg :=NewVectorPolar (_deee /2,_eba +_dge /2);_bff :=NewVectorPolar (_gb ,_eba );_dgg :=NewVectorPolar (_ggg ,_eba +_dge /2);_fgbe :=_caf .AddVector (_dgg );_gcf :=_bff .Add (_gfg .Flip ());_afa :=_fgbe .AddVector (_gcf );_fbde :=_gfg .Scale (2).Flip ().Add (_gcf .Flip ());_fec :=_afa .AddVector (_fbde );_ccd :=_caf .AddVector (NewVectorPolar (_fa ,_eba -_dge /2));_cfb :=NewPath ();_cfb =_cfb .AppendPoint (_dbad .GetPointNumber (1));_cfb =_cfb .AppendPoint (_caf );_cfb =_cfb .AppendPoint (_fgbe );_cfb =_cfb .AppendPoint (_afa );_cfb =_cfb .AppendPoint (_fec );_cfb =_cfb .AppendPoint (_ccd );_cfb =_cfb .AppendPoint (_dbad .GetPointNumber (4));_dbad =_cfb ;};if _eeb ==LineEndingStyleArrow {_cab :=_dbad .GetPointNumber (1);_bab :=_dbad .GetPointNumber (_dbad .Length ());_bba :=NewVectorPolar (_fa /2,_eba +_dge +_dge /2);_cfc :=_cab .AddVector (_bba );_fd :=NewVectorPolar (_gb ,_eba ).Add (NewVectorPolar (_deee /2,_eba +_dge /2));_ddd :=_cfc .AddVector (_fd );_faf :=NewVectorPolar (_ggg ,_eba -_dge /2);_dae :=_ddd .AddVector (_faf );_edc :=NewVectorPolar (_gb ,_eba );_ace :=_bab .AddVector (_edc );_fcg :=NewVectorPolar (_ggg ,_eba +_dge +_dge /2);_gbe :=_ace .AddVector (_fcg );_gba :=_cfc ;_aca :=NewPath ();_aca =_aca .AppendPoint (_cfc );_aca =_aca .AppendPoint (_ddd );_aca =_aca .AppendPoint (_dae );for _ ,_gd :=range _dbad .Points [1:len (_dbad .Points )-1]{_aca =_aca .AppendPoint (_gd );};_aca =_aca .AppendPoint (_ace );_aca =_aca .AppendPoint (_gbe );_aca =_aca .AppendPoint (_gba );_dbad =_aca ;};_fcbg :=_df .NewContentCreator ();_fcbg .Add_q ().Add_rg (_fgf .LineColor .R (),_fgf .LineColor .G (),_fgf .LineColor .B ());if len (gsName )> 1{_fcbg .Add_gs (_f .PdfObjectName (gsName ));};_dbad =_dbad .Offset (_fgf .X1 ,_fgf .Y1 );_fggb :=_dbad .GetBoundingBox ();DrawPathWithCreator (_dbad ,_fcbg );if _fgf .LineStyle ==LineStyleDashed {_fcbg .Add_d ([]int64 {1,1},0).Add_S ().Add_f ().Add_Q ();}else {_fcbg .Add_f ().Add_Q ();};return _fcbg .Bytes (),_fggb .ToPdfRectangle (),nil ;};

// Add shifts the coordinates of the point with dx, dy and returns the result.
func (_gga Point )Add (dx ,dy float64 )Point {_gga .X +=dx ;_gga .Y +=dy ;return _gga };

// Path consists of straight line connections between each point defined in an array of points.
type Path struct{Points []Point ;};

// Rotate rotates the vector by the specified angle.
func (_aeb Vector )Rotate (phi float64 )Vector {_gbg :=_aeb .Magnitude ();_dfd :=_aeb .GetPolarAngle ();return NewVectorPolar (_gbg ,_dfd +phi );};

// FlipY flips the sign of the Dy component of the vector.
func (_fece Vector )FlipY ()Vector {_fece .Dy =-_fece .Dy ;return _fece };

// GetPointNumber returns the path point at the index specified by number.
// The index is 1-based.
func (_ae Path )GetPointNumber (number int )Point {if number < 1||number > len (_ae .Points ){return Point {};};return _ae .Points [number -1];};

// NewVectorPolar returns a new vector calculated from the specified
// magnitude and angle.
func NewVectorPolar (length float64 ,theta float64 )Vector {_deg :=Vector {};_deg .Dx =length *_e .Cos (theta );_deg .Dy =length *_e .Sin (theta );return _deg ;};

// Vector represents a two-dimensional vector.
type Vector struct{Dx float64 ;Dy float64 ;};

// AppendPoint adds the specified point to the path.
func (_bc Path )AppendPoint (point Point )Path {_bc .Points =append (_bc .Points ,point );return _bc };

// NewVector returns a new vector with the direction specified by dx and dy.
func NewVector (dx ,dy float64 )Vector {_baa :=Vector {};_baa .Dx =dx ;_baa .Dy =dy ;return _baa };const (LineStyleSolid LineStyle =0;LineStyleDashed LineStyle =1;);

// BasicLine defines a line between point 1 (X1,Y1) and point 2 (X2,Y2). The line has a specified width, color and opacity.
type BasicLine struct{X1 float64 ;Y1 float64 ;X2 float64 ;Y2 float64 ;LineColor *_a .PdfColorDeviceRGB ;Opacity float64 ;LineWidth float64 ;LineStyle LineStyle ;};

// Add adds the specified vector to the current one and returns the result.
func (_ddf Vector )Add (other Vector )Vector {_ddf .Dx +=other .Dx ;_ddf .Dy +=other .Dy ;return _ddf };

// GetBoundingBox returns the bounding box of the Bezier path.
func (_ac CubicBezierPath )GetBoundingBox ()Rectangle {_af :=Rectangle {};_ffa :=0.0;_ce :=0.0;_bb :=0.0;_bd :=0.0;for _dca ,_dba :=range _ac .Curves {_ba :=_dba .GetBounds ();if _dca ==0{_ffa =_ba .Llx ;_ce =_ba .Urx ;_bb =_ba .Lly ;_bd =_ba .Ury ;continue ;};if _ba .Llx < _ffa {_ffa =_ba .Llx ;};if _ba .Urx > _ce {_ce =_ba .Urx ;};if _ba .Lly < _bb {_bb =_ba .Lly ;};if _ba .Ury > _bd {_bd =_ba .Ury ;};};_af .X =_ffa ;_af .Y =_bb ;_af .Width =_ce -_ffa ;_af .Height =_bd -_bb ;return _af ;};

// Draw draws the polyline. A graphics state name can be specified for
// setting the polyline properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polyline
// bounding box.
func (_gea Polyline )Draw (gsName string )([]byte ,*_a .PdfRectangle ,error ){if _gea .LineColor ==nil {_gea .LineColor =_a .NewPdfColorDeviceRGB (0,0,0);};_eac :=NewPath ();for _ ,_bbg :=range _gea .Points {_eac =_eac .AppendPoint (_bbg );};_gec :=_df .NewContentCreator ();_gec .Add_q ();_gec .Add_RG (_gea .LineColor .R (),_gea .LineColor .G (),_gea .LineColor .B ());_gec .Add_w (_gea .LineWidth );if len (gsName )> 1{_gec .Add_gs (_f .PdfObjectName (gsName ));};DrawPathWithCreator (_eac ,_gec );_gec .Add_S ();_gec .Add_Q ();return _gec .Bytes (),_eac .GetBoundingBox ().ToPdfRectangle (),nil ;};const (LineEndingStyleNone LineEndingStyle =0;LineEndingStyleArrow LineEndingStyle =1;LineEndingStyleButt LineEndingStyle =2;);

// Rectangle is a shape with a specified Width and Height and a lower left corner at (X,Y) that can be
// drawn to a PDF content stream.  The rectangle can optionally have a border and a filling color.
// The Width/Height includes the border (if any specified), i.e. is positioned inside.
type Rectangle struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;FillEnabled bool ;FillColor *_a .PdfColorDeviceRGB ;BorderEnabled bool ;BorderWidth float64 ;BorderColor *_a .PdfColorDeviceRGB ;Opacity float64 ;};

// Polygon is a multi-point shape that can be drawn to a PDF content stream.
type Polygon struct{Points [][]Point ;FillEnabled bool ;FillColor *_a .PdfColorDeviceRGB ;BorderEnabled bool ;BorderColor *_a .PdfColorDeviceRGB ;BorderWidth float64 ;};

// AddOffsetXY adds X,Y offset to all points on a curve.
func (_fg CubicBezierCurve )AddOffsetXY (offX ,offY float64 )CubicBezierCurve {_fg .P0 .X +=offX ;_fg .P1 .X +=offX ;_fg .P2 .X +=offX ;_fg .P3 .X +=offX ;_fg .P0 .Y +=offY ;_fg .P1 .Y +=offY ;_fg .P2 .Y +=offY ;_fg .P3 .Y +=offY ;return _fg ;};

// Scale scales the vector by the specified factor.
func (_bgb Vector )Scale (factor float64 )Vector {_bbge :=_bgb .Magnitude ();_ecc :=_bgb .GetPolarAngle ();_bgb .Dx =factor *_bbge *_e .Cos (_ecc );_bgb .Dy =factor *_bbge *_e .Sin (_ecc );return _bgb ;};

// AppendCurve appends the specified Bezier curve to the path.
func (_ag CubicBezierPath )AppendCurve (curve CubicBezierCurve )CubicBezierPath {_ag .Curves =append (_ag .Curves ,curve );return _ag ;};

// Draw draws the rectangle. Can specify a graphics state (gsName) for setting opacity etc.
// Otherwise leave empty (""). Returns the content stream as a byte array, bounding box and an error on failure.
func (_fgb Rectangle )Draw (gsName string )([]byte ,*_a .PdfRectangle ,error ){_bgf :=NewPath ();_bgf =_bgf .AppendPoint (NewPoint (0,0));_bgf =_bgf .AppendPoint (NewPoint (0,_fgb .Height ));_bgf =_bgf .AppendPoint (NewPoint (_fgb .Width ,_fgb .Height ));_bgf =_bgf .AppendPoint (NewPoint (_fgb .Width ,0));_bgf =_bgf .AppendPoint (NewPoint (0,0));if _fgb .X !=0||_fgb .Y !=0{_bgf =_bgf .Offset (_fgb .X ,_fgb .Y );};_feg :=_df .NewContentCreator ();_feg .Add_q ();if _fgb .FillEnabled {_feg .Add_rg (_fgb .FillColor .R (),_fgb .FillColor .G (),_fgb .FillColor .B ());};if _fgb .BorderEnabled {_feg .Add_RG (_fgb .BorderColor .R (),_fgb .BorderColor .G (),_fgb .BorderColor .B ());_feg .Add_w (_fgb .BorderWidth );};if len (gsName )> 1{_feg .Add_gs (_f .PdfObjectName (gsName ));};DrawPathWithCreator (_bgf ,_feg );_feg .Add_h ();if _fgb .FillEnabled &&_fgb .BorderEnabled {_feg .Add_B ();}else if _fgb .FillEnabled {_feg .Add_f ();}else if _fgb .BorderEnabled {_feg .Add_S ();};_feg .Add_Q ();return _feg .Bytes (),_bgf .GetBoundingBox ().ToPdfRectangle (),nil ;};

// Circle represents a circle shape with fill and border properties that can be drawn to a PDF content stream.
type Circle struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;FillEnabled bool ;FillColor *_a .PdfColorDeviceRGB ;BorderEnabled bool ;BorderWidth float64 ;BorderColor *_a .PdfColorDeviceRGB ;Opacity float64 ;};

// Draw draws the basic line to PDF. Generates the content stream which can be used in page contents or appearance
// stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error if
// one occurred.
func (_ef BasicLine )Draw (gsName string )([]byte ,*_a .PdfRectangle ,error ){_fgc :=_ef .LineWidth ;_baf :=NewPath ();_baf =_baf .AppendPoint (NewPoint (_ef .X1 ,_ef .Y1 ));_baf =_baf .AppendPoint (NewPoint (_ef .X2 ,_ef .Y2 ));_ebac :=_df .NewContentCreator ();_fgd :=_baf .GetBoundingBox ();DrawPathWithCreator (_baf ,_ebac );if _ef .LineStyle ==LineStyleDashed {_ebac .Add_d ([]int64 {1,1},0);};_ebac .Add_RG (_ef .LineColor .R (),_ef .LineColor .G (),_ef .LineColor .B ()).Add_w (_fgc ).Add_S ().Add_Q ();return _ebac .Bytes (),_fgd .ToPdfRectangle (),nil ;};

// Magnitude returns the magnitude of the vector.
func (_acac Vector )Magnitude ()float64 {return _e .Sqrt (_e .Pow (_acac .Dx ,2.0)+_e .Pow (_acac .Dy ,2.0));};

// GetBoundingBox returns the bounding box of the path.
func (_fb Path )GetBoundingBox ()BoundingBox {_ecf :=BoundingBox {};_eed :=0.0;_cf :=0.0;_fgg :=0.0;_bg :=0.0;for _cd ,_ea :=range _fb .Points {if _cd ==0{_eed =_ea .X ;_cf =_ea .X ;_fgg =_ea .Y ;_bg =_ea .Y ;continue ;};if _ea .X < _eed {_eed =_ea .X ;};if _ea .X > _cf {_cf =_ea .X ;};if _ea .Y < _fgg {_fgg =_ea .Y ;};if _ea .Y > _bg {_bg =_ea .Y ;};};_ecf .X =_eed ;_ecf .Y =_fgg ;_ecf .Width =_cf -_eed ;_ecf .Height =_bg -_fgg ;return _ecf ;};

// NewVectorBetween returns a new vector with the direction specified by
// the subtraction of point a from point b (b-a).
func NewVectorBetween (a Point ,b Point )Vector {_fgcc :=Vector {};_fgcc .Dx =b .X -a .X ;_fgcc .Dy =b .Y -a .Y ;return _fgcc ;};

// NewPoint returns a new point with the coordinates x, y.
func NewPoint (x ,y float64 )Point {return Point {X :x ,Y :y }};

// DrawBezierPathWithCreator makes the bezier path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawBezierPathWithCreator (bpath CubicBezierPath ,creator *_df .ContentCreator ){for _gfd ,_gfc :=range bpath .Curves {if _gfd ==0{creator .Add_m (_gfc .P0 .X ,_gfc .P0 .Y );};creator .Add_c (_gfc .P1 .X ,_gfc .P1 .Y ,_gfc .P2 .X ,_gfc .P2 .Y ,_gfc .P3 .X ,_gfc .P3 .Y );};};

// Copy returns a clone of the path.
func (_ccg Path )Copy ()Path {_gg :=Path {};_gg .Points =[]Point {};for _ ,_bcf :=range _ccg .Points {_gg .Points =append (_gg .Points ,_bcf );};return _gg ;};

// BoundingBox represents the smallest rectangular area that encapsulates an object.
type BoundingBox struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;};

// Offset shifts the path with the specified offsets.
func (_bf Path )Offset (offX ,offY float64 )Path {for _fga ,_bde :=range _bf .Points {_bf .Points [_fga ]=_bde .Add (offX ,offY );};return _bf ;};

// FlipX flips the sign of the Dx component of the vector.
func (_dcg Vector )FlipX ()Vector {_dcg .Dx =-_dcg .Dx ;return _dcg };

// Offset shifts the Bezier path with the specified offsets.
func (_ec CubicBezierPath )Offset (offX ,offY float64 )CubicBezierPath {for _da ,_fc :=range _ec .Curves {_ec .Curves [_da ]=_fc .AddOffsetXY (offX ,offY );};return _ec ;};

// CubicBezierCurve is defined by:
// R(t) = P0*(1-t)^3 + P1*3*t*(1-t)^2 + P2*3*t^2*(1-t) + P3*t^3
// where P0 is the current point, P1, P2 control points and P3 the final point.
type CubicBezierCurve struct{P0 Point ;P1 Point ;P2 Point ;P3 Point ;};

// Draw draws the composite Bezier curve. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array and
// the curve bounding box.
func (_bda PolyBezierCurve )Draw (gsName string )([]byte ,*_a .PdfRectangle ,error ){if _bda .BorderColor ==nil {_bda .BorderColor =_a .NewPdfColorDeviceRGB (0,0,0);};_aea :=NewCubicBezierPath ();for _ ,_bcb :=range _bda .Curves {_aea =_aea .AppendCurve (_bcb );};_dd :=_df .NewContentCreator ();_dd .Add_q ();_bda .FillEnabled =_bda .FillEnabled &&_bda .FillColor !=nil ;if _bda .FillEnabled {_dd .Add_rg (_bda .FillColor .R (),_bda .FillColor .G (),_bda .FillColor .B ());};_dd .Add_RG (_bda .BorderColor .R (),_bda .BorderColor .G (),_bda .BorderColor .B ());_dd .Add_w (_bda .BorderWidth );if len (gsName )> 1{_dd .Add_gs (_f .PdfObjectName (gsName ));};for _ ,_fce :=range _aea .Curves {_dd .Add_m (_fce .P0 .X ,_fce .P0 .Y );_dd .Add_c (_fce .P1 .X ,_fce .P1 .Y ,_fce .P2 .X ,_fce .P2 .Y ,_fce .P3 .X ,_fce .P3 .Y );};if _bda .FillEnabled {_dd .Add_h ();_dd .Add_B ();}else {_dd .Add_S ();};_dd .Add_Q ();return _dd .Bytes (),_aea .GetBoundingBox ().ToPdfRectangle (),nil ;};

// Line defines a line shape between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none (regular line),
// or arrows at either end.  The line also has a specified width, color and opacity.
type Line struct{X1 float64 ;Y1 float64 ;X2 float64 ;Y2 float64 ;LineColor *_a .PdfColorDeviceRGB ;Opacity float64 ;LineWidth float64 ;LineEndingStyle1 LineEndingStyle ;LineEndingStyle2 LineEndingStyle ;LineStyle LineStyle ;};

// CubicBezierPath represents a collection of cubic Bezier curves.
type CubicBezierPath struct{Curves []CubicBezierCurve ;};

// NewPath returns a new empty path.
func NewPath ()Path {return Path {}};