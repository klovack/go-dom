package godom

// DocTypeDeclaration is a type of a doctype
type DocTypeDeclaration = string

var (
	// None DocType for empty doctype
	None DocTypeDeclaration
	// HTML5 DocType
	HTML5 DocTypeDeclaration = "html"
	// HTML401S DocType
	HTML401S DocTypeDeclaration = `HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd"`
	// HTML401T DocType
	HTML401T DocTypeDeclaration = `HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd`
	// HTML401F DocType
	HTML401F DocTypeDeclaration = `HTML PUBLIC "-//W3C//DTD HTML 4.01 Frameset//EN" "http://www.w3.org/TR/html4/frameset.dtd"`
	// XHTML10S DocType
	XHTML10S DocTypeDeclaration = `html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd"`
	// XHTML10T DocType
	XHTML10T DocTypeDeclaration = `html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd"`
	// XHTML10F DocType
	XHTML10F DocTypeDeclaration = `html PUBLIC "-//W3C//DTD XHTML 1.0 Frameset//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd"`
	// XHTML11 DocType
	XHTML11 DocTypeDeclaration = `html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd"`
)
