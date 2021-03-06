//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		nvd.nist.gov/schema/cce_0.1.xsd
package go_Cce01

//	CCE is at an early phase of adoption.  This schema is a work in progress and is far from
//	final.  Additional work with using CCEs in a practical setting is required.

import (
	scap_core "github.com/linuxisnotunix/Vulnerobot/modules/models/xsd/nvd.nist.gov/schema/scap-core_0.1.xsd_go"
	xsdt "github.com/metaleap/go-xsd/types"
)

type XsdGoPkgHasElems_ValuesequencecceParameterTypeschema_Value_XsdtString_ struct {
	Values []xsdt.String `xml:"http://scap.nist.gov/schema/cce/0.1 value"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_ValuesequencecceParameterTypeschema_Value_XsdtString_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_ValuesequencecceParameterTypeschema_Value_XsdtString_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_ValuesequencecceParameterTypeschema_Value_XsdtString_ instance.
func (me *XsdGoPkgHasElems_ValuesequencecceParameterTypeschema_Value_XsdtString_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_ValuesequencecceParameterTypeschema_Value_XsdtString_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

//	TODO: What does this identify?
type XsdGoPkgHasAttr_Identifier_XsdtToken_ struct {
	//	TODO: What does this identify?
	Identifier xsdt.Token `xml:"http://scap.nist.gov/schema/cce/0.1 identifier,attr"`
}

//	TODO: should this be an enumeration?
type XsdGoPkgHasAttr_Operator_XsdtToken_ struct {
	//	TODO: should this be an enumeration?
	Operator xsdt.Token `xml:"http://scap.nist.gov/schema/cce/0.1 operator,attr"`
}

type TcceParameterType struct {
	XsdGoPkgHasElems_ValuesequencecceParameterTypeschema_Value_XsdtString_

	//	TODO: What does this identify?
	XsdGoPkgHasAttr_Identifier_XsdtToken_

	//	TODO: should this be an enumeration?
	XsdGoPkgHasAttr_Operator_XsdtToken_
}

//	If the WalkHandlers.TcceParameterType function is not nil (ie. was set by outside code), calls it with this TcceParameterType instance as the single argument. Then calls the Walk() method on 1/3 embed(s) and 0/0 field(s) belonging to this TcceParameterType instance.
func (me *TcceParameterType) Walk() (err error) {
	if fn := WalkHandlers.TcceParameterType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.XsdGoPkgHasElems_ValuesequencecceParameterTypeschema_Value_XsdtString_.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElem_DefinitionsequencecceTypeschema_Definition_XsdtString_ struct {
	Definition xsdt.String `xml:"http://scap.nist.gov/schema/cce/0.1 definition"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_DefinitionsequencecceTypeschema_Definition_XsdtString_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_DefinitionsequencecceTypeschema_Definition_XsdtString_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElem_DefinitionsequencecceTypeschema_Definition_XsdtString_ instance.
func (me *XsdGoPkgHasElem_DefinitionsequencecceTypeschema_Definition_XsdtString_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_DefinitionsequencecceTypeschema_Definition_XsdtString_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElem_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_ struct {
	References scap_core.TreferenceType `xml:"http://scap.nist.gov/schema/cce/0.1 references"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElem_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_ instance.
func (me *XsdGoPkgHasElem_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

//	The format for a CCE name is CCE-NNNNNNNNNNN, where NNNNNNNNNNN is a sequence number.
type TcceNamePatternType xsdt.Token

//	Since TcceNamePatternType is just a simple String type, this merely sets the current value from the specified string.
func (me *TcceNamePatternType) Set(s string) { (*xsdt.Token)(me).Set(s) }

//	Since TcceNamePatternType is just a simple String type, this merely returns the current string value.
func (me TcceNamePatternType) String() string { return xsdt.Token(me).String() }

//	This convenience method just performs a simple type conversion to TcceNamePatternType's alias type xsdt.Token.
func (me TcceNamePatternType) ToXsdtToken() xsdt.Token { return xsdt.Token(me) }

type XsdGoPkgHasElems_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_ struct {
	TechnicalMechanismses []xsdt.String `xml:"http://scap.nist.gov/schema/cce/0.1 technical-mechanisms"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_ instance.
func (me *XsdGoPkgHasElems_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_ParametersequencecceTypeschema_Parameter_TcceParameterType_ struct {
	Parameters []*TcceParameterType `xml:"http://scap.nist.gov/schema/cce/0.1 parameter"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_ParametersequencecceTypeschema_Parameter_TcceParameterType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_ParametersequencecceTypeschema_Parameter_TcceParameterType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_ParametersequencecceTypeschema_Parameter_TcceParameterType_ instance.
func (me *XsdGoPkgHasElems_ParametersequencecceTypeschema_Parameter_TcceParameterType_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_ParametersequencecceTypeschema_Parameter_TcceParameterType_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		for _, x := range me.Parameters {
			if err = x.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_ struct {
	Referenceses []scap_core.TreferenceType `xml:"http://scap.nist.gov/schema/cce/0.1 references"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_ instance.
func (me *XsdGoPkgHasElems_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElem_ValuesequencecceParameterTypeschema_Value_XsdtString_ struct {
	Value xsdt.String `xml:"http://scap.nist.gov/schema/cce/0.1 value"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_ValuesequencecceParameterTypeschema_Value_XsdtString_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_ValuesequencecceParameterTypeschema_Value_XsdtString_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElem_ValuesequencecceParameterTypeschema_Value_XsdtString_ instance.
func (me *XsdGoPkgHasElem_ValuesequencecceParameterTypeschema_Value_XsdtString_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_ValuesequencecceParameterTypeschema_Value_XsdtString_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasCdata struct {
	XsdGoPkgCDATA string `xml:",chardata"`
}

//	If the WalkHandlers.XsdGoPkgHasCdata function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasCdata instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasCdata instance.
func (me *XsdGoPkgHasCdata) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasCdata; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasAttr_Id_TcceNamePatternType_ struct {
	Id TcceNamePatternType `xml:"http://scap.nist.gov/schema/cce/0.1 id,attr"`
}

type XsdGoPkgHasElem_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_ struct {
	TechnicalMechanisms xsdt.String `xml:"http://scap.nist.gov/schema/cce/0.1 technical-mechanisms"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElem_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_ instance.
func (me *XsdGoPkgHasElem_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type TcceType struct {
	XsdGoPkgHasElem_DefinitionsequencecceTypeschema_Definition_XsdtString_

	XsdGoPkgHasElems_ParametersequencecceTypeschema_Parameter_TcceParameterType_

	XsdGoPkgHasElems_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_

	XsdGoPkgHasElems_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_

	XsdGoPkgHasAttr_Id_TcceNamePatternType_
}

//	If the WalkHandlers.TcceType function is not nil (ie. was set by outside code), calls it with this TcceType instance as the single argument. Then calls the Walk() method on 4/5 embed(s) and 0/0 field(s) belonging to this TcceType instance.
func (me *TcceType) Walk() (err error) {
	if fn := WalkHandlers.TcceType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.XsdGoPkgHasElems_ParametersequencecceTypeschema_Parameter_TcceParameterType_.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if err = me.XsdGoPkgHasElems_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if err = me.XsdGoPkgHasElems_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if err = me.XsdGoPkgHasElem_DefinitionsequencecceTypeschema_Definition_XsdtString_.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_DefinitionsequencecceTypeschema_Definition_XsdtString_ struct {
	Definitions []xsdt.String `xml:"http://scap.nist.gov/schema/cce/0.1 definition"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_DefinitionsequencecceTypeschema_Definition_XsdtString_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_DefinitionsequencecceTypeschema_Definition_XsdtString_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_DefinitionsequencecceTypeschema_Definition_XsdtString_ instance.
func (me *XsdGoPkgHasElems_DefinitionsequencecceTypeschema_Definition_XsdtString_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_DefinitionsequencecceTypeschema_Definition_XsdtString_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElem_ParametersequencecceTypeschema_Parameter_TcceParameterType_ struct {
	Parameter *TcceParameterType `xml:"http://scap.nist.gov/schema/cce/0.1 parameter"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_ParametersequencecceTypeschema_Parameter_TcceParameterType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_ParametersequencecceTypeschema_Parameter_TcceParameterType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElem_ParametersequencecceTypeschema_Parameter_TcceParameterType_ instance.
func (me *XsdGoPkgHasElem_ParametersequencecceTypeschema_Parameter_TcceParameterType_) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_ParametersequencecceTypeschema_Parameter_TcceParameterType_; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.Parameter.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

var (
	//	Set this to false to break a Walk() immediately as soon as the first error is returned by a custom handler function.
	//	If true, Walk() proceeds and accumulates all errors in the WalkErrors slice.
	WalkContinueOnError = true
	//	Contains all errors accumulated during Walk()s. If you're using this, you need to reset this yourself as needed prior to a fresh Walk().
	WalkErrors []error
	//	Your custom error-handling function, if required.
	WalkOnError func(error)
	//	Provides 13 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
	//	If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
	WalkHandlers = &XsdGoPkgWalkHandlers{}
)

//	Provides 13 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
//	If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
type XsdGoPkgWalkHandlers struct {
	XsdGoPkgHasElems_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_       func(*XsdGoPkgHasElems_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_, bool) error
	XsdGoPkgHasCdata                                                                          func(*XsdGoPkgHasCdata, bool) error
	XsdGoPkgHasElems_DefinitionsequencecceTypeschema_Definition_XsdtString_                   func(*XsdGoPkgHasElems_DefinitionsequencecceTypeschema_Definition_XsdtString_, bool) error
	XsdGoPkgHasElem_ParametersequencecceTypeschema_Parameter_TcceParameterType_               func(*XsdGoPkgHasElem_ParametersequencecceTypeschema_Parameter_TcceParameterType_, bool) error
	TcceParameterType                                                                         func(*TcceParameterType, bool) error
	XsdGoPkgHasElem_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_        func(*XsdGoPkgHasElem_ReferencessequencecceTypeschema_References_ScapCoreTreferenceType_, bool) error
	XsdGoPkgHasElems_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_ func(*XsdGoPkgHasElems_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_, bool) error
	XsdGoPkgHasElems_ParametersequencecceTypeschema_Parameter_TcceParameterType_              func(*XsdGoPkgHasElems_ParametersequencecceTypeschema_Parameter_TcceParameterType_, bool) error
	XsdGoPkgHasElem_ValuesequencecceParameterTypeschema_Value_XsdtString_                     func(*XsdGoPkgHasElem_ValuesequencecceParameterTypeschema_Value_XsdtString_, bool) error
	XsdGoPkgHasElem_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_  func(*XsdGoPkgHasElem_TechnicalMechanismssequencecceTypeschema_TechnicalMechanisms_XsdtString_, bool) error
	TcceType                                                                                  func(*TcceType, bool) error
	XsdGoPkgHasElems_ValuesequencecceParameterTypeschema_Value_XsdtString_                    func(*XsdGoPkgHasElems_ValuesequencecceParameterTypeschema_Value_XsdtString_, bool) error
	XsdGoPkgHasElem_DefinitionsequencecceTypeschema_Definition_XsdtString_                    func(*XsdGoPkgHasElem_DefinitionsequencecceTypeschema_Definition_XsdtString_, bool) error
}
