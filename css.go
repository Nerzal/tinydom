package tinydom

import "syscall/js"

type CSS struct {
	js.Value
}

func (s *CSS) CssText() string {
	return s.Get("cssText").String()
}

func (s *CSS) Length() int {
	return s.Get("length").Int()
}

func (s *CSS) AlignContent() string {
	return s.Get("alignContent").String()
}

func (s *CSS) SetAlignContent(v string) {
	s.Set("alignContent", v)
}

func (s *CSS) AlignItems() string {
	return s.Get("alignItems").String()
}

func (s *CSS) SetAlignItems(v string) {
	s.Set("alignItems", v)
}

func (s *CSS) AlignSelf() string {
	return s.Get("alignSelf").String()
}

func (s *CSS) SetAlignSelf(v string) {
	s.Set("alignSelf", v)
}

func (s *CSS) Animation() string {
	return s.Get("animation").String()
}

func (s *CSS) SetAnimation(v string) {
	s.Set("animation", v)
}

func (s *CSS) AnimationDelay() string {
	return s.Get("animationDelay").String()
}

func (s *CSS) SetAnimationDelay(v string) {
	s.Set("animationDelay", v)
}

func (s *CSS) AnimationDirection() string {
	return s.Get("animationDirection").String()
}

func (s *CSS) SetAnimationDirection(v string) {
	s.Set("animationDirection", v)
}

func (s *CSS) AnimationDuration() string {
	return s.Get("animationDuration").String()
}

func (s *CSS) SetAnimationDuration(v string) {
	s.Set("animationDuration", v)
}

func (s *CSS) AnimationFillMode() string {
	return s.Get("animationFillMode").String()
}

func (s *CSS) SetAnimationFillMode(v string) {
	s.Set("animationFillMode", v)
}

func (s *CSS) AnimationIterationCount() string {
	return s.Get("animationIterationCount").String()
}

func (s *CSS) SetAnimationIterationCount(v string) {
	s.Set("animationIterationCount", v)
}

func (s *CSS) AnimationName() string {
	return s.Get("animationName").String()
}

func (s *CSS) SetAnimationName(v string) {
	s.Set("animationName", v)
}

func (s *CSS) AnimationTimingFunction() string {
	return s.Get("animationTimingFunction").String()
}

func (s *CSS) SetAnimationTimingFunction(v string) {
	s.Set("animationTimingFunction", v)
}

func (s *CSS) AnimationPlayState() string {
	return s.Get("animationPlayState").String()
}

func (s *CSS) SetAnimationPlayState(v string) {
	s.Set("animationPlayState", v)
}

func (s *CSS) Background() string {
	return s.Get("background").String()
}

func (s *CSS) SetBackground(v string) {
	s.Set("background", v)
}

func (s *CSS) BackgroundAttachment() string {
	return s.Get("backgroundAttachment").String()
}

func (s *CSS) SetBackgroundAttachment(v string) {
	s.Set("backgroundAttachment", v)
}

func (s *CSS) BackgroundColor() string {
	return s.Get("backgroundColor").String()
}

func (s *CSS) SetBackgroundColor(v string) {
	s.Set("backgroundColor", v)
}

func (s *CSS) BackgroundImage() string {
	return s.Get("backgroundImage").String()
}

func (s *CSS) SetBackgroundImage(v string) {
	s.Set("backgroundImage", v)
}

func (s *CSS) BackgroundPosition() string {
	return s.Get("backgroundPosition").String()
}

func (s *CSS) SetBackgroundPosition(v string) {
	s.Set("backgroundPosition", v)
}

func (s *CSS) BackgroundRepeat() string {
	return s.Get("backgroundRepeat").String()
}

func (s *CSS) SetBackgroundRepeat(v string) {
	s.Set("backgroundRepeat", v)
}

func (s *CSS) BackgroundClip() string {
	return s.Get("backgroundClip").String()
}

func (s *CSS) SetBackgroundClip(v string) {
	s.Set("backgroundClip", v)
}

func (s *CSS) BackgroundOrigin() string {
	return s.Get("backgroundOrigin").String()
}

func (s *CSS) SetBackgroundOrigin(v string) {
	s.Set("backgroundOrigin", v)
}

func (s *CSS) BackgroundSize() string {
	return s.Get("backgroundSize").String()
}

func (s *CSS) SetBackgroundSize(v string) {
	s.Set("backgroundSize", v)
}

func (s *CSS) BackfaceVisibility() string {
	return s.Get("backfaceVisibility").String()
}

func (s *CSS) SetBackfaceVisibility(v string) {
	s.Set("backfaceVisibility", v)
}

func (s *CSS) Border() string {
	return s.Get("border").String()
}

func (s *CSS) SetBorder(v string) {
	s.Set("border", v)
}

func (s *CSS) BorderBottom() string {
	return s.Get("borderBottom").String()
}

func (s *CSS) SetBorderBottom(v string) {
	s.Set("borderBottom", v)
}

func (s *CSS) BorderBottomColor() string {
	return s.Get("borderBottomColor").String()
}

func (s *CSS) SetBorderBottomColor(v string) {
	s.Set("borderBottomColor", v)
}

func (s *CSS) BorderBottomLeftRadius() string {
	return s.Get("borderBottomLeftRadius").String()
}

func (s *CSS) SetBorderBottomLeftRadius(v string) {
	s.Set("borderBottomLeftRadius", v)
}

func (s *CSS) BorderBottomRightRadius() string {
	return s.Get("borderBottomRightRadius").String()
}

func (s *CSS) SetBorderBottomRightRadius(v string) {
	s.Set("borderBottomRightRadius", v)
}

func (s *CSS) BorderBottomStyle() string {
	return s.Get("borderBottomStyle").String()
}

func (s *CSS) SetBorderBottomStyle(v string) {
	s.Set("borderBottomStyle", v)
}

func (s *CSS) BorderBottomWidth() string {
	return s.Get("borderBottomWidth").String()
}

func (s *CSS) SetBorderBottomWidth(v string) {
	s.Set("borderBottomWidth", v)
}

func (s *CSS) BorderCollapse() string {
	return s.Get("borderCollapse").String()
}

func (s *CSS) SetBorderCollapse(v string) {
	s.Set("borderCollapse", v)
}

func (s *CSS) BorderColor() string {
	return s.Get("borderColor").String()
}

func (s *CSS) SetBorderColor(v string) {
	s.Set("borderColor", v)
}

func (s *CSS) BorderImage() string {
	return s.Get("borderImage").String()
}

func (s *CSS) SetBorderImage(v string) {
	s.Set("borderImage", v)
}

func (s *CSS) BorderImageOutset() string {
	return s.Get("borderImageOutset").String()
}

func (s *CSS) SetBorderImageOutset(v string) {
	s.Set("borderImageOutset", v)
}

func (s *CSS) BorderImageRepeat() string {
	return s.Get("borderImageRepeat").String()
}

func (s *CSS) SetBorderImageRepeat(v string) {
	s.Set("borderImageRepeat", v)
}

func (s *CSS) BorderImageSlice() string {
	return s.Get("borderImageSlice").String()
}

func (s *CSS) SetBorderImageSlice(v string) {
	s.Set("borderImageSlice", v)
}

func (s *CSS) BorderImageSource() string {
	return s.Get("borderImageSource").String()
}

func (s *CSS) SetBorderImageSource(v string) {
	s.Set("borderImageSource", v)
}

func (s *CSS) BorderImageWidth() string {
	return s.Get("borderImageWidth").String()
}

func (s *CSS) SetBorderImageWidth(v string) {
	s.Set("borderImageWidth", v)
}

func (s *CSS) BorderLeft() string {
	return s.Get("borderLeft").String()
}

func (s *CSS) SetBorderLeft(v string) {
	s.Set("borderLeft", v)
}

func (s *CSS) BorderLeftColor() string {
	return s.Get("borderLeftColor").String()
}

func (s *CSS) SetBorderLeftColor(v string) {
	s.Set("borderLeftColor", v)
}

func (s *CSS) BorderLeftStyle() string {
	return s.Get("borderLeftStyle").String()
}

func (s *CSS) SetBorderLeftStyle(v string) {
	s.Set("borderLeftStyle", v)
}

func (s *CSS) BorderLeftWidth() string {
	return s.Get("borderLeftWidth").String()
}

func (s *CSS) SetBorderLeftWidth(v string) {
	s.Set("borderLeftWidth", v)
}

func (s *CSS) BorderRadius() string {
	return s.Get("borderRadius").String()
}

func (s *CSS) SetBorderRadius(v string) {
	s.Set("borderRadius", v)
}

func (s *CSS) BorderRight() string {
	return s.Get("borderRight").String()
}

func (s *CSS) SetBorderRight(v string) {
	s.Set("borderRight", v)
}

func (s *CSS) BorderRightColor() string {
	return s.Get("borderRightColor").String()
}

func (s *CSS) SetBorderRightColor(v string) {
	s.Set("borderRightColor", v)
}

func (s *CSS) BorderRightStyle() string {
	return s.Get("borderRightStyle").String()
}

func (s *CSS) SetBorderRightStyle(v string) {
	s.Set("borderRightStyle", v)
}

func (s *CSS) BorderRightWidth() string {
	return s.Get("borderRightWidth").String()
}

func (s *CSS) SetBorderRightWidth(v string) {
	s.Set("borderRightWidth", v)
}

func (s *CSS) BorderSpacing() string {
	return s.Get("borderSpacing").String()
}

func (s *CSS) SetBorderSpacing(v string) {
	s.Set("borderSpacing", v)
}

func (s *CSS) BorderStyle() string {
	return s.Get("borderStyle").String()
}

func (s *CSS) SetBorderStyle(v string) {
	s.Set("borderStyle", v)
}

func (s *CSS) BorderTop() string {
	return s.Get("borderTop").String()
}

func (s *CSS) SetBorderTop(v string) {
	s.Set("borderTop", v)
}

func (s *CSS) BorderTopColor() string {
	return s.Get("borderTopColor").String()
}

func (s *CSS) SetBorderTopColor(v string) {
	s.Set("borderTopColor", v)
}

func (s *CSS) BorderTopLeftRadius() string {
	return s.Get("borderTopLeftRadius").String()
}

func (s *CSS) SetBorderTopLeftRadius(v string) {
	s.Set("borderTopLeftRadius", v)
}

func (s *CSS) BorderTopRightRadius() string {
	return s.Get("borderTopRightRadius").String()
}

func (s *CSS) SetBorderTopRightRadius(v string) {
	s.Set("borderTopRightRadius", v)
}

func (s *CSS) BorderTopStyle() string {
	return s.Get("borderTopStyle").String()
}

func (s *CSS) SetBorderTopStyle(v string) {
	s.Set("borderTopStyle", v)
}

func (s *CSS) BorderTopWidth() string {
	return s.Get("borderTopWidth").String()
}

func (s *CSS) SetBorderTopWidth(v string) {
	s.Set("borderTopWidth", v)
}

func (s *CSS) BorderWidth() string {
	return s.Get("borderWidth").String()
}

func (s *CSS) SetBorderWidth(v string) {
	s.Set("borderWidth", v)
}

func (s *CSS) Bottom() string {
	return s.Get("bottom").String()
}

func (s *CSS) SetBottom(v string) {
	s.Set("bottom", v)
}

func (s *CSS) BoxShadow() string {
	return s.Get("boxShadow").String()
}

func (s *CSS) SetBoxShadow(v string) {
	s.Set("boxShadow", v)
}

func (s *CSS) BoxSizing() string {
	return s.Get("boxSizing").String()
}

func (s *CSS) SetBoxSizing(v string) {
	s.Set("boxSizing", v)
}

func (s *CSS) CaptionSide() string {
	return s.Get("captionSide").String()
}

func (s *CSS) SetCaptionSide(v string) {
	s.Set("captionSide", v)
}

func (s *CSS) Clear() string {
	return s.Get("clear").String()
}

func (s *CSS) SetClear(v string) {
	s.Set("clear", v)
}

func (s *CSS) Clip() string {
	return s.Get("clip").String()
}

func (s *CSS) SetClip(v string) {
	s.Set("clip", v)
}

func (s *CSS) Color() string {
	return s.Get("color").String()
}

func (s *CSS) SetColor(v string) {
	s.Set("color", v)
}

func (s *CSS) ColumnCount() string {
	return s.Get("columnCount").String()
}

func (s *CSS) SetColumnCount(v string) {
	s.Set("columnCount", v)
}

func (s *CSS) ColumnFill() string {
	return s.Get("columnFill").String()
}

func (s *CSS) SetColumnFill(v string) {
	s.Set("columnFill", v)
}

func (s *CSS) ColumnGap() string {
	return s.Get("columnGap").String()
}

func (s *CSS) SetColumnGap(v string) {
	s.Set("columnGap", v)
}

func (s *CSS) ColumnRule() string {
	return s.Get("columnRule").String()
}

func (s *CSS) SetColumnRule(v string) {
	s.Set("columnRule", v)
}

func (s *CSS) ColumnRuleColor() string {
	return s.Get("columnRuleColor").String()
}

func (s *CSS) SetColumnRuleColor(v string) {
	s.Set("columnRuleColor", v)
}

func (s *CSS) ColumnRuleStyle() string {
	return s.Get("columnRuleStyle").String()
}

func (s *CSS) SetColumnRuleStyle(v string) {
	s.Set("columnRuleStyle", v)
}

func (s *CSS) ColumnRuleWidth() string {
	return s.Get("columnRuleWidth").String()
}

func (s *CSS) SetColumnRuleWidth(v string) {
	s.Set("columnRuleWidth", v)
}

func (s *CSS) Columns() string {
	return s.Get("columns").String()
}

func (s *CSS) SetColumns(v string) {
	s.Set("columns", v)
}

func (s *CSS) ColumnSpan() string {
	return s.Get("columnSpan").String()
}

func (s *CSS) SetColumnSpan(v string) {
	s.Set("columnSpan", v)
}

func (s *CSS) ColumnWidth() string {
	return s.Get("columnWidth").String()
}

func (s *CSS) SetColumnWidth(v string) {
	s.Set("columnWidth", v)
}

func (s *CSS) CounterIncrement() string {
	return s.Get("counterIncrement").String()
}

func (s *CSS) SetCounterIncrement(v string) {
	s.Set("counterIncrement", v)
}

func (s *CSS) CounterReset() string {
	return s.Get("counterReset").String()
}

func (s *CSS) SetCounterReset(v string) {
	s.Set("counterReset", v)
}

func (s *CSS) Cursor() string {
	return s.Get("cursor").String()
}

func (s *CSS) SetCursor(v string) {
	s.Set("cursor", v)
}

func (s *CSS) Direction() string {
	return s.Get("direction").String()
}

func (s *CSS) SetDirection(v string) {
	s.Set("direction", v)
}

func (s *CSS) Display() string {
	return s.Get("display").String()
}

func (s *CSS) SetDisplay(v string) {
	s.Set("display", v)
}

func (s *CSS) EmptyCells() string {
	return s.Get("emptyCells").String()
}

func (s *CSS) SetEmptyCells(v string) {
	s.Set("emptyCells", v)
}

func (s *CSS) Filter() string {
	return s.Get("filter").String()
}

func (s *CSS) SetFilter(v string) {
	s.Set("filter", v)
}

func (s *CSS) Flex() string {
	return s.Get("flex").String()
}

func (s *CSS) SetFlex(v string) {
	s.Set("flex", v)
}

func (s *CSS) FlexBasis() string {
	return s.Get("flexBasis").String()
}

func (s *CSS) SetFlexBasis(v string) {
	s.Set("flexBasis", v)
}

func (s *CSS) FlexDirection() string {
	return s.Get("flexDirection").String()
}

func (s *CSS) SetFlexDirection(v string) {
	s.Set("flexDirection", v)
}

func (s *CSS) FlexFlow() string {
	return s.Get("flexFlow").String()
}

func (s *CSS) SetFlexFlow(v string) {
	s.Set("flexFlow", v)
}

func (s *CSS) FlexGrow() string {
	return s.Get("flexGrow").String()
}

func (s *CSS) SetFlexGrow(v string) {
	s.Set("flexGrow", v)
}

func (s *CSS) FlexShrink() string {
	return s.Get("flexShrink").String()
}

func (s *CSS) SetFlexShrink(v string) {
	s.Set("flexShrink", v)
}

func (s *CSS) FlexWrap() string {
	return s.Get("flexWrap").String()
}

func (s *CSS) SetFlexWrap(v string) {
	s.Set("flexWrap", v)
}

func (s *CSS) CssFloat() string {
	return s.Get("cssFloat").String()
}

func (s *CSS) SetCssFloat(v string) {
	s.Set("cssFloat", v)
}

func (s *CSS) Font() string {
	return s.Get("font").String()
}

func (s *CSS) SetFont(v string) {
	s.Set("font", v)
}

func (s *CSS) FontFamily() string {
	return s.Get("fontFamily").String()
}

func (s *CSS) SetFontFamily(v string) {
	s.Set("fontFamily", v)
}

func (s *CSS) FontSize() string {
	return s.Get("fontSize").String()
}

func (s *CSS) SetFontSize(v string) {
	s.Set("fontSize", v)
}

func (s *CSS) FontStyle() string {
	return s.Get("fontStyle").String()
}

func (s *CSS) SetFontStyle(v string) {
	s.Set("fontStyle", v)
}

func (s *CSS) FontVariant() string {
	return s.Get("fontVariant").String()
}

func (s *CSS) SetFontVariant(v string) {
	s.Set("fontVariant", v)
}

func (s *CSS) FontWeight() string {
	return s.Get("fontWeight").String()
}

func (s *CSS) SetFontWeight(v string) {
	s.Set("fontWeight", v)
}

func (s *CSS) FontSizeAdjust() string {
	return s.Get("fontSizeAdjust").String()
}

func (s *CSS) SetFontSizeAdjust(v string) {
	s.Set("fontSizeAdjust", v)
}

func (s *CSS) Height() string {
	return s.Get("height").String()
}

func (s *CSS) SetHeight(v string) {
	s.Set("height", v)
}

func (s *CSS) JustifyContent() string {
	return s.Get("justifyContent").String()
}

func (s *CSS) SetJustifyContent(v string) {
	s.Set("justifyContent", v)
}

func (s *CSS) Left() string {
	return s.Get("left").String()
}

func (s *CSS) SetLeft(v string) {
	s.Set("left", v)
}

func (s *CSS) LetterSpacing() string {
	return s.Get("letterSpacing").String()
}

func (s *CSS) SetLetterSpacing(v string) {
	s.Set("letterSpacing", v)
}

func (s *CSS) LineHeight() string {
	return s.Get("lineHeight").String()
}

func (s *CSS) SetLineHeight(v string) {
	s.Set("lineHeight", v)
}

func (s *CSS) ListStyle() string {
	return s.Get("listStyle").String()
}

func (s *CSS) SetListStyle(v string) {
	s.Set("listStyle", v)
}

func (s *CSS) ListStyleImage() string {
	return s.Get("listStyleImage").String()
}

func (s *CSS) SetListStyleImage(v string) {
	s.Set("listStyleImage", v)
}

func (s *CSS) ListStylePosition() string {
	return s.Get("listStylePosition").String()
}

func (s *CSS) SetListStylePosition(v string) {
	s.Set("listStylePosition", v)
}

func (s *CSS) ListStyleType() string {
	return s.Get("listStyleType").String()
}

func (s *CSS) SetListStyleType(v string) {
	s.Set("listStyleType", v)
}

func (s *CSS) Margin() string {
	return s.Get("margin").String()
}

func (s *CSS) SetMargin(v string) {
	s.Set("margin", v)
}

func (s *CSS) MarginBottom() string {
	return s.Get("marginBottom").String()
}

func (s *CSS) SetMarginBottom(v string) {
	s.Set("marginBottom", v)
}

func (s *CSS) MarginLeft() string {
	return s.Get("marginLeft").String()
}

func (s *CSS) SetMarginLeft(v string) {
	s.Set("marginLeft", v)
}

func (s *CSS) MarginRight() string {
	return s.Get("marginRight").String()
}

func (s *CSS) SetMarginRight(v string) {
	s.Set("marginRight", v)
}

func (s *CSS) MarginTop() string {
	return s.Get("marginTop").String()
}

func (s *CSS) SetMarginTop(v string) {
	s.Set("marginTop", v)
}

func (s *CSS) MaxHeight() string {
	return s.Get("maxHeight").String()
}

func (s *CSS) SetMaxHeight(v string) {
	s.Set("maxHeight", v)
}

func (s *CSS) MaxWidth() string {
	return s.Get("maxWidth").String()
}

func (s *CSS) SetMaxWidth(v string) {
	s.Set("maxWidth", v)
}

func (s *CSS) MinHeight() string {
	return s.Get("minHeight").String()
}

func (s *CSS) SetMinHeight(v string) {
	s.Set("minHeight", v)
}

func (s *CSS) MinWidth() string {
	return s.Get("minWidth").String()
}

func (s *CSS) SetMinWidth(v string) {
	s.Set("minWidth", v)
}

func (s *CSS) Opacity() string {
	return s.Get("opacity").String()
}

func (s *CSS) SetOpacity(v string) {
	s.Set("opacity", v)
}

func (s *CSS) Order() string {
	return s.Get("order").String()
}

func (s *CSS) SetOrder(v string) {
	s.Set("order", v)
}

func (s *CSS) Orphans() string {
	return s.Get("orphans").String()
}

func (s *CSS) SetOrphans(v string) {
	s.Set("orphans", v)
}

func (s *CSS) Outline() string {
	return s.Get("outline").String()
}

func (s *CSS) SetOutline(v string) {
	s.Set("outline", v)
}

func (s *CSS) OutlineColor() string {
	return s.Get("outlineColor").String()
}

func (s *CSS) SetOutlineColor(v string) {
	s.Set("outlineColor", v)
}

func (s *CSS) OutlineOffset() string {
	return s.Get("outlineOffset").String()
}

func (s *CSS) SetOutlineOffset(v string) {
	s.Set("outlineOffset", v)
}

func (s *CSS) OutlineStyle() string {
	return s.Get("outlineStyle").String()
}

func (s *CSS) SetOutlineStyle(v string) {
	s.Set("outlineStyle", v)
}

func (s *CSS) OutlineWidth() string {
	return s.Get("outlineWidth").String()
}

func (s *CSS) SetOutlineWidth(v string) {
	s.Set("outlineWidth", v)
}

func (s *CSS) Overflow() string {
	return s.Get("overflow").String()
}

func (s *CSS) SetOverflow(v string) {
	s.Set("overflow", v)
}

func (s *CSS) OverflowX() string {
	return s.Get("overflowX").String()
}

func (s *CSS) SetOverflowX(v string) {
	s.Set("overflowX", v)
}

func (s *CSS) OverflowY() string {
	return s.Get("overflowY").String()
}

func (s *CSS) SetOverflowY(v string) {
	s.Set("overflowY", v)
}

func (s *CSS) Padding() string {
	return s.Get("padding").String()
}

func (s *CSS) SetPadding(v string) {
	s.Set("padding", v)
}

func (s *CSS) PaddingBottom() string {
	return s.Get("paddingBottom").String()
}

func (s *CSS) SetPaddingBottom(v string) {
	s.Set("paddingBottom", v)
}

func (s *CSS) PaddingLeft() string {
	return s.Get("paddingLeft").String()
}

func (s *CSS) SetPaddingLeft(v string) {
	s.Set("paddingLeft", v)
}

func (s *CSS) PaddingRight() string {
	return s.Get("paddingRight").String()
}

func (s *CSS) SetPaddingRight(v string) {
	s.Set("paddingRight", v)
}

func (s *CSS) PaddingTop() string {
	return s.Get("paddingTop").String()
}

func (s *CSS) SetPaddingTop(v string) {
	s.Set("paddingTop", v)
}

func (s *CSS) PageBreakAfter() string {
	return s.Get("pageBreakAfter").String()
}

func (s *CSS) SetPageBreakAfter(v string) {
	s.Set("pageBreakAfter", v)
}

func (s *CSS) PageBreakBefore() string {
	return s.Get("pageBreakBefore").String()
}

func (s *CSS) SetPageBreakBefore(v string) {
	s.Set("pageBreakBefore", v)
}

func (s *CSS) PageBreakInside() string {
	return s.Get("pageBreakInside").String()
}

func (s *CSS) SetPageBreakInside(v string) {
	s.Set("pageBreakInside", v)
}

func (s *CSS) Perspective() string {
	return s.Get("perspective").String()
}

func (s *CSS) SetPerspective(v string) {
	s.Set("perspective", v)
}

func (s *CSS) PerspectiveOrigin() string {
	return s.Get("perspectiveOrigin").String()
}

func (s *CSS) SetPerspectiveOrigin(v string) {
	s.Set("perspectiveOrigin", v)
}

func (s *CSS) Position() string {
	return s.Get("position").String()
}

func (s *CSS) SetPosition(v string) {
	s.Set("position", v)
}

func (s *CSS) Quotes() string {
	return s.Get("quotes").String()
}

func (s *CSS) SetQuotes(v string) {
	s.Set("quotes", v)
}

func (s *CSS) Resize() string {
	return s.Get("resize").String()
}

func (s *CSS) SetResize(v string) {
	s.Set("resize", v)
}

func (s *CSS) Right() string {
	return s.Get("right").String()
}

func (s *CSS) SetRight(v string) {
	s.Set("right", v)
}

func (s *CSS) TableLayout() string {
	return s.Get("tableLayout").String()
}

func (s *CSS) SetTableLayout(v string) {
	s.Set("tableLayout", v)
}

func (s *CSS) TabSize() string {
	return s.Get("tabSize").String()
}

func (s *CSS) SetTabSize(v string) {
	s.Set("tabSize", v)
}

func (s *CSS) TextAlign() string {
	return s.Get("textAlign").String()
}

func (s *CSS) SetTextAlign(v string) {
	s.Set("textAlign", v)
}

func (s *CSS) TextAlignLast() string {
	return s.Get("textAlignLast").String()
}

func (s *CSS) SetTextAlignLast(v string) {
	s.Set("textAlignLast", v)
}

func (s *CSS) TextDecoration() string {
	return s.Get("textDecoration").String()
}

func (s *CSS) SetTextDecoration(v string) {
	s.Set("textDecoration", v)
}

func (s *CSS) TextDecorationColor() string {
	return s.Get("textDecorationColor").String()
}

func (s *CSS) SetTextDecorationColor(v string) {
	s.Set("textDecorationColor", v)
}

func (s *CSS) TextDecorationLine() string {
	return s.Get("textDecorationLine").String()
}

func (s *CSS) SetTextDecorationLine(v string) {
	s.Set("textDecorationLine", v)
}

func (s *CSS) TextDecorationStyle() string {
	return s.Get("textDecorationStyle").String()
}

func (s *CSS) SetTextDecorationStyle(v string) {
	s.Set("textDecorationStyle", v)
}

func (s *CSS) TextIndent() string {
	return s.Get("textIndent").String()
}

func (s *CSS) SetTextIndent(v string) {
	s.Set("textIndent", v)
}

func (s *CSS) TextOverflow() string {
	return s.Get("textOverflow").String()
}

func (s *CSS) SetTextOverflow(v string) {
	s.Set("textOverflow", v)
}

func (s *CSS) TextShadow() string {
	return s.Get("textShadow").String()
}

func (s *CSS) SetTextShadow(v string) {
	s.Set("textShadow", v)
}

func (s *CSS) TextTransform() string {
	return s.Get("textTransform").String()
}

func (s *CSS) SetTextTransform(v string) {
	s.Set("textTransform", v)
}

func (s *CSS) Top() string {
	return s.Get("top").String()
}

func (s *CSS) SetTop(v string) {
	s.Set("top", v)
}

func (s *CSS) Transform() string {
	return s.Get("transform").String()
}

func (s *CSS) SetTransform(v string) {
	s.Set("transform", v)
}

func (s *CSS) TransformOrigin() string {
	return s.Get("transformOrigin").String()
}

func (s *CSS) SetTransformOrigin(v string) {
	s.Set("transformOrigin", v)
}

func (s *CSS) TransformStyle() string {
	return s.Get("transformStyle").String()
}

func (s *CSS) SetTransformStyle(v string) {
	s.Set("transformStyle", v)
}

func (s *CSS) Transition() string {
	return s.Get("transition").String()
}

func (s *CSS) SetTransition(v string) {
	s.Set("transition", v)
}

func (s *CSS) TransitionProperty() string {
	return s.Get("transitionProperty").String()
}

func (s *CSS) SetTransitionProperty(v string) {
	s.Set("transitionProperty", v)
}

func (s *CSS) TransitionDuration() string {
	return s.Get("transitionDuration").String()
}

func (s *CSS) SetTransitionDuration(v string) {
	s.Set("transitionDuration", v)
}

func (s *CSS) TransitionTimingFunction() string {
	return s.Get("transitionTimingFunction").String()
}

func (s *CSS) SetTransitionTimingFunction(v string) {
	s.Set("transitionTimingFunction", v)
}

func (s *CSS) TransitionDelay() string {
	return s.Get("transitionDelay").String()
}

func (s *CSS) SetTransitionDelay(v string) {
	s.Set("transitionDelay", v)
}

func (s *CSS) UnicodeBidi() string {
	return s.Get("unicodeBidi").String()
}

func (s *CSS) SetUnicodeBidi(v string) {
	s.Set("unicodeBidi", v)
}

func (s *CSS) UserSelect() string {
	return s.Get("userSelect").String()
}

func (s *CSS) SetUserSelect(v string) {
	s.Set("userSelect", v)
}

func (s *CSS) VerticalAlign() string {
	return s.Get("verticalAlign").String()
}

func (s *CSS) SetVerticalAlign(v string) {
	s.Set("verticalAlign", v)
}

func (s *CSS) Visibility() string {
	return s.Get("visibility").String()
}

func (s *CSS) SetVisibility(v string) {
	s.Set("visibility", v)
}

func (s *CSS) WhiteSpace() string {
	return s.Get("whiteSpace").String()
}

func (s *CSS) SetWhiteSpace(v string) {
	s.Set("whiteSpace", v)
}

func (s *CSS) Width() string {
	return s.Get("width").String()
}

func (s *CSS) SetWidth(v string) {
	s.Set("width", v)
}

func (s *CSS) WordBreak() string {
	return s.Get("wordBreak").String()
}

func (s *CSS) SetWordBreak(v string) {
	s.Set("wordBreak", v)
}

func (s *CSS) WordSpacing() string {
	return s.Get("wordSpacing").String()
}

func (s *CSS) SetWordSpacing(v string) {
	s.Set("wordSpacing", v)
}

func (s *CSS) WordWrap() string {
	return s.Get("wordWrap").String()
}

func (s *CSS) SetWordWrap(v string) {
	s.Set("wordWrap", v)
}

func (s *CSS) Widows() string {
	return s.Get("widows").String()
}

func (s *CSS) SetWidows(v string) {
	s.Set("widows", v)
}

func (s *CSS) ZIndex() string {
	return s.Get("zIndex").String()
}

func (s *CSS) SetZIndex(v string) {
	s.Set("zIndex", v)
}
