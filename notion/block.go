package notion

import (
	"encoding/json"
	"fmt"
	"time"
)

type BlockList struct {
	Object     string      `json:"object"`
	Results    []Block     `json:"results"`
	NextCursor interface{} `json:"next_cursor"`
	HasMore    bool        `json:"has_more"`
}

func (b *BlockList) UnmarshalJSON(data []byte) error {
	type Alias BlockList
	a := &struct {
		Results []interface{} `json:"results"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	for _, val := range a.Results {
		msg := val.(map[string]interface{})
		switch msg["type"] {
		case "to_do":
			jsonObj, err := json.Marshal(val)
			if err != nil {
				return err
			}

			var block ToDoBlock
			err = json.Unmarshal(jsonObj, &block)
			if err != nil {
				return err
			} else {
				fmt.Println(block.Text)
			}

		}
		//switch val["type"] {
		////case "paragraph":
		////	var block ParagraphBlock
		////	err := json.Unmarshal(jsonObj, &block)
		////	if err != nil {
		////		return err
		////	} else {
		////		b.Results = append(b.Results, block)
		////	}
		////case "heading_1":
		////	var block HeadingOneBlock
		////	err := json.Unmarshal(jsonObj, &block)
		////	if err != nil {
		////		return err
		////	} else {
		////		b.Results = append(b.Results, block)
		////	}
		//case "to_do":
		//	fmt.Println(val)
		//	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		//	jsonObj, err := json.Marshal(val)
		//	if err != nil {
		//		return err
		//	}
		//
		//	var block ToDoBlock
		//	err = json.Unmarshal(jsonObj, &block)
		//	if err != nil {
		//		return err
		//	} else {
		//		fmt.Println(block)
		//		b.Results = append(b.Results, block)
		//	}
		//}
	}

	return nil
}

type Block interface {
	PlainText() string
}

type BlockBase struct {
	Object         string    `json:"object"`
	Id             string    `json:"id"`
	CreatedTime    time.Time `json:"created_time"`
	LastEditedTime time.Time `json:"last_edited_time"`
	HasChildren    bool      `json:"has_children"`
	Type           string    `json:"type"`
}

type ParagraphBlock struct {
	BlockBase

	Text     []RichText  `json:"text"`
	Children []BlockBase `json:"children"`
}

type HeadingOneBlock struct {
	Text []RichText `json:"text"`
}

type HeadingTwoBlock struct {
	Text []RichText `json:"text"`
}

type HeadingThreeBlock struct {
	Text []RichText `json:"text"`
}

type BulletedListItemBlock struct {
	Text     []RichText  `json:"text"`
	Children []BlockBase `json:"children"`
}

type NumberedListItemBlock struct {
	Text     []RichText  `json:"text"`
	Children []BlockBase `json:"children"`
}

type ToDoBlock struct {
	BlockBase

	Text     []RichText  `json:"text"`
	Children []BlockBase `json:"children"`
	Checked  bool        `json:"checked"`
}

func (t ToDoBlock) PlainText() string {
	return t.Text[0].PlainText
}

type ToggleBlock struct {
	Text     []RichText  `json:"text"`
	Children []BlockBase `json:"children"`
}

type ChildPageBlock struct {
	Text  []RichText `json:"text"`
	Title string     `json:"title"`
}

type UnsupportedBlock struct {
}

type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

type RichTextBase struct {
	PlainText   string      `json:"plain_text"`
	Href        string      `json:"href"`
	Type        string      `json:"type"`
	Annotations Annotations `json:"annotations"`
}

type RichText struct {
	RichTextBase

	Text struct {
		Content string `json:"content"`
		Link    struct {
			Type string `json:"type"`
			Url  string `json:"url"`
		} `json:"link"`
	} `json:"text"`
}
