package ip2asn

import "errors"

//This package can return net.DNSError
//It can also return any error as follows

//ErrBadIPLength indicates an IP address not length 4 (v4) or 16 (v6).
var ErrBadIPLength = errors.New("bad IP length")

//ErrTooManyAnswers indicates a DNS query returned more than one answer.
var ErrTooManyAnswers = errors.New("too many answers")

//ErrNotImplimented indicates a function is not yet implimented
var ErrNotImplimented = errors.New("not implimented")

//ErrInvalidASNRecord indicates an ASN record could not be parsed
var ErrInvalidASNRecord = errors.New("invalid asn record, cannot parse")

//ErrInvalidPrefixString indicates a prefix string could not be parsed
var ErrInvalidPrefixString = errors.New("invalid prefix string")

//ErrInvalidAnswerFormat indicates the DNS answer could not be parsed
var ErrInvalidAnswerFormat = errors.New("invalid answer format, cannot parse")

//ErrInvalidPrefixRecord indicates a prefix record could not be parsed
var ErrInvalidPrefixRecord = errors.New("invalid prefix record")

//ErrInvalidRelation indicates a Relation value received was invalid
var ErrInvalidRelation = errors.New("invlid relation")
