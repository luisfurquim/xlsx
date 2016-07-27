package xlsx



type TimeCell struct {
   Type       string `xml:"value-type,attr"`
   Val        string `xml:"p"`
   DisplayVal string `xml:"time-value,attr"`
}

type DateCell struct {
   Type       string `xml:"value-type,attr"`
   Val        string `xml:"p"`
   DisplayVal string `xml:"date-value,attr"`
}

type Cell struct {
   Type       string `xml:"t,attr"`
   Val        string `xml:"v"`
}

type Row struct {
   Cells []Cell `xml:"c"`
}

type Data struct {
   Rows  []Row `xml:"row"`
}

type Document struct {
//   Data  Data `xml:"worksheet>sheetdata"`
   Data  Data `xml:"sheetData"`
}

type Document2 struct {
   Row   []Cell `xml:"row"`
}

