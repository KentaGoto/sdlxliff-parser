package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Structure to parse XML in sdlxliff
type Xliff struct {
	XMLName xml.Name `xml:"xliff"`
	Text    string   `xml:",chardata"`
	Sdl     string   `xml:"sdl,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Version string   `xml:"version,attr"`
	File    struct {
		Text           string `xml:",chardata"`
		Original       string `xml:"original,attr"`
		Datatype       string `xml:"datatype,attr"`
		SourceLanguage string `xml:"source-language,attr"`
		TargetLanguage string `xml:"target-language,attr"`
		Header         struct {
			Text      string `xml:",chardata"`
			Reference struct {
				Text         string `xml:",chardata"`
				InternalFile struct {
					Text string `xml:",chardata"`
					Form string `xml:"form,attr"`
				} `xml:"internal-file"`
			} `xml:"reference"`
			RefFiles struct {
				Text    string `xml:",chardata"`
				RefFile struct {
					Text        string `xml:",chardata"`
					Uid         string `xml:"uid,attr"`
					ID          string `xml:"id,attr"`
					Name        string `xml:"name,attr"`
					OPath       string `xml:"o-path,attr"`
					Date        string `xml:"date,attr"`
					ExpectedUse string `xml:"expected-use,attr"`
				} `xml:"ref-file"`
			} `xml:"ref-files"`
			FileInfo struct {
				Text  string `xml:",chardata"`
				Xmlns string `xml:"xmlns,attr"`
				Value []struct {
					Text string `xml:",chardata"`
					Key  string `xml:"key,attr"`
				} `xml:"value"`
				SniffInfo struct {
					Text               string `xml:",chardata"`
					DetectedSourceLang struct {
						Text           string `xml:",chardata"`
						DetectionLevel string `xml:"detection-level,attr"`
						Lang           string `xml:"lang,attr"`
					} `xml:"detected-source-lang"`
				} `xml:"sniff-info"`
			} `xml:"file-info"`
			FiletypeInfo struct {
				Text       string `xml:",chardata"`
				FiletypeID string `xml:"filetype-id"`
			} `xml:"filetype-info"`
			FmtDefs struct {
				Text   string `xml:",chardata"`
				Xmlns  string `xml:"xmlns,attr"`
				FmtDef []struct {
					Text  string `xml:",chardata"`
					ID    string `xml:"id,attr"`
					Value []struct {
						Text string `xml:",chardata"`
						Key  string `xml:"key,attr"`
					} `xml:"value"`
				} `xml:"fmt-def"`
			} `xml:"fmt-defs"`
			CxtDefs struct {
				Text   string `xml:",chardata"`
				Xmlns  string `xml:"xmlns,attr"`
				CxtDef []struct {
					Text    string `xml:",chardata"`
					ID      string `xml:"id,attr"`
					Type    string `xml:"type,attr"`
					Purpose string `xml:"purpose,attr"`
					Name    string `xml:"name,attr"`
					Code    string `xml:"code,attr"`
					Descr   string `xml:"descr,attr"`
					Color   string `xml:"color,attr"`
					Fmt     struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
					} `xml:"fmt"`
					Props struct {
						Text  string `xml:",chardata"`
						Value []struct {
							Text string `xml:",chardata"`
							Key  string `xml:"key,attr"`
						} `xml:"value"`
					} `xml:"props"`
				} `xml:"cxt-def"`
			} `xml:"cxt-defs"`
			NodeDefs struct {
				Text    string `xml:",chardata"`
				Xmlns   string `xml:"xmlns,attr"`
				NodeDef []struct {
					Text      string `xml:",chardata"`
					ID        string `xml:"id,attr"`
					Parent    string `xml:"parent,attr"`
					ForceName string `xml:"force-name,attr"`
					Cxt       struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
					} `xml:"cxt"`
				} `xml:"node-def"`
			} `xml:"node-defs"`
			TagDefs struct {
				Text  string `xml:",chardata"`
				Xmlns string `xml:"xmlns,attr"`
				Tag   []struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Bpt  struct {
						Text    string `xml:",chardata"`
						Name    string `xml:"name,attr"`
						CanHide string `xml:"can-hide,attr"`
						WordEnd string `xml:"word-end,attr"`
						SegHint string `xml:"seg-hint,attr"`
						Sub     struct {
							Text string `xml:",chardata"`
							Xid  string `xml:"xid,attr"`
						} `xml:"sub"`
					} `xml:"bpt"`
					BptProps struct {
						Text  string `xml:",chardata"`
						Value []struct {
							Text string `xml:",chardata"`
							Key  string `xml:"key,attr"`
						} `xml:"value"`
					} `xml:"bpt-props"`
					Ept struct {
						Text    string `xml:",chardata"`
						Name    string `xml:"name,attr"`
						CanHide string `xml:"can-hide,attr"`
						WordEnd string `xml:"word-end,attr"`
					} `xml:"ept"`
					Fmt struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
					} `xml:"fmt"`
					Ph struct {
						Text    string `xml:",chardata"`
						Name    string `xml:"name,attr"`
						WordEnd string `xml:"word-end,attr"`
						SegHint string `xml:"seg-hint,attr"`
					} `xml:"ph"`
					Props struct {
						Text  string `xml:",chardata"`
						Value []struct {
							Text string `xml:",chardata"`
							Key  string `xml:"key,attr"`
						} `xml:"value"`
					} `xml:"props"`
				} `xml:"tag"`
			} `xml:"tag-defs"`
		} `xml:"header"`
		Body struct {
			Text  string `xml:",chardata"`
			Group []struct {
				Text string `xml:",chardata"`
				Cxts struct {
					Text string `xml:",chardata"`
					Cxt  []struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
					} `xml:"cxt"`
					Node struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
					} `xml:"node"`
				} `xml:"cxts"`
				TransUnit struct {
					Text      string `xml:",chardata"`
					ID        string `xml:"id,attr"`
					Translate string `xml:"translate,attr"`
					Source    struct {
						Text string `xml:",chardata"`
						G    []struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Xid  string `xml:"xid,attr"`
							X    struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
							} `xml:"x"`
						} `xml:"g"`
						X []struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"x"`
					} `xml:"source"`
					SegSource struct {
						Text string `xml:",chardata"`
						G    struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Mrk  []struct {
								Text  string `xml:",chardata"`
								Mtype string `xml:"mtype,attr"`
								Mid   string `xml:"mid,attr"`
							} `xml:"mrk"`
						} `xml:"g"`
						Mrk []struct {
							Text  string `xml:",chardata"`
							Mtype string `xml:"mtype,attr"`
							Mid   string `xml:"mid,attr"`
							X     []struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
							} `xml:"x"`
							G []struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Xid  string `xml:"xid,attr"`
								X    struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
								} `xml:"x"`
							} `xml:"g"`
						} `xml:"mrk"`
						X []struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"x"`
					} `xml:"seg-source"`
					Target struct {
						Text string `xml:",chardata"`
						G    struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Mrk  []struct {
								Text  string `xml:",chardata"`
								Mtype string `xml:"mtype,attr"`
								Mid   string `xml:"mid,attr"`
							} `xml:"mrk"`
						} `xml:"g"`
						Mrk []struct {
							Text  string `xml:",chardata"`
							Mtype string `xml:"mtype,attr"`
							Mid   string `xml:"mid,attr"`
							X     []struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
							} `xml:"x"`
							G struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
							} `xml:"g"`
						} `xml:"mrk"`
						X []struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"x"`
					} `xml:"target"`
					SegDefs struct {
						Text string `xml:",chardata"`
						Seg  []struct {
							Text         string `xml:",chardata"`
							ID           string `xml:"id,attr"`
							Conf         string `xml:"conf,attr"`
							Origin       string `xml:"origin,attr"`
							OriginSystem string `xml:"origin-system,attr"`
							Percent      string `xml:"percent,attr"`
							StructMatch  string `xml:"struct-match,attr"`
							TextMatch    string `xml:"text-match,attr"`
							Value        []struct {
								Text string `xml:",chardata"`
								Key  string `xml:"key,attr"`
							} `xml:"value"`
						} `xml:"seg"`
					} `xml:"seg-defs"`
				} `xml:"trans-unit"`
			} `xml:"group"`
		} `xml:"body"`
	} `xml:"file"`
}

func main() {
	// Takes a target file as an argument.
	fname := os.Args[1]

	// Load sdlxliff file.
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	// Parses sdlxliff and stores data in a structure.
	var sdl Xliff
	err = xml.Unmarshal(data, &sdl)
	if err != nil {
		fmt.Println("XML Unmarshal error", err)
		return
	}

	fmt.Println()
	fmt.Println("Original:", sdl.File.Original)
	fmt.Println("Datatype:", sdl.File.Datatype)
	fmt.Println("SourceLanguage:", sdl.File.SourceLanguage)
	fmt.Println("TargetLanguage:", sdl.File.TargetLanguage)
	fmt.Println()
	if len(sdl.File.Body.Group) > 0 {
		for _, gr := range sdl.File.Body.Group {
			// src
			for _, mr := range gr.TransUnit.SegSource.Mrk {
				if len(mr.Text) > 0 {
					fmt.Println("[Source]:", mr.Text)
				} else {
					for _, mrg := range mr.G {
						fmt.Println("[Source]:", mrg.Text)
					}
				}
			}
			// target
			for _, mr := range gr.TransUnit.Target.Mrk {
				if len(mr.Text) > 0 {
					fmt.Println("[Target]:", mr.Text)
				} else {
					fmt.Println("[Target]:", mr.G.Text)
				}
			}
		}
	} else {
		fmt.Println("Empty.")
	}
}
