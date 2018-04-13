package surveyMonkeyCode

import (
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"fmt"
	"strconv"
)

func SetSurveyDetails(jsonstr string, jsonSR string, activityOutput string) string {
		//set metadata
		questions := gjson.Get(jsonstr, "surveydetails.pages.0.questions")
		for _, que := range questions.Array() {
			//fmt.Println("que= ",que.String())
			queIndex := gjson.Get(activityOutput, "survey.questions.#").String()
			fmt.Println("queIndex= "+queIndex)
			//set heading
			activityOutput_tmp, _ := sjson.Set(activityOutput, "survey.questions."+queIndex+".title", gjson.Get(que.String(), "headings.0.heading").String())
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".id", gjson.Get(que.String(), "id").String())
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".validation", gjson.Get(que.String(), "validation").String())
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".position", gjson.Get(que.String(), "position").String())
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".subtype", gjson.Get(que.String(), "subtype").String())
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".family", gjson.Get(que.String(), "family").String())
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".type", gjson.Get(que.String(), "required.type").String())
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".visible", gjson.Get(que.String(), "visible").String())
			//set answers-rows
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.rows.0.id", "")
			rows := gjson.Get(que.String(), "answers.rows")
			for r, row := range rows.Array() {
					activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.rows."+strconv.Itoa(r)+".visible", gjson.Get(row.String(), "visible").String())
					activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.rows."+strconv.Itoa(r)+".text", gjson.Get(row.String(), "text").String())
					activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.rows."+strconv.Itoa(r)+".position", gjson.Get(row.String(), "position").String())
					activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.rows."+strconv.Itoa(r)+".id", gjson.Get(row.String(), "id").String())
			}
			//set answers-other
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.other.id", gjson.Get(que.String(), "answers.other.id").String())
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.other.visible", gjson.Get(que.String(), "answers.other.visible").String())
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.other.is_answer_choice", gjson.Get(que.String(), "answers.other.is_answer_choice").String())
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.other.apply_all_rows", gjson.Get(que.String(), "answers.other.apply_all_rows").String())
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.other.text", gjson.Get(que.String(), "answers.other.text").String())
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.other.position", gjson.Get(que.String(), "answers.other.position").String())
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.other.num_chars", gjson.Get(que.String(), "answers.other.num_chars").String())
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.other.error_text", gjson.Get(que.String(), "answers.other.error_text").String())
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.other.num_lines", gjson.Get(que.String(), "answers.other.num_lines").String())
			//set answer-choices
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.choices.0.id", "")
			choice := gjson.Get(que.String(), "answers.choices")
			for c, ch := range choice.Array() {
					activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.choices."+strconv.Itoa(c)+".visible", gjson.Get(ch.String(), "visible").String())
					activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.choices."+strconv.Itoa(c)+".text", gjson.Get(ch.String(), "text").String())
					activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.choices."+strconv.Itoa(c)+".position", gjson.Get(ch.String(), "position").String())
					activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.choices."+strconv.Itoa(c)+".is_na", gjson.Get(ch.String(), "is_na").String())
					activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.choices."+strconv.Itoa(c)+".weight", gjson.Get(ch.String(), "weight").String())
					activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.choices."+strconv.Itoa(c)+".description", gjson.Get(ch.String(), "description").String())
					activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".answers.choices."+strconv.Itoa(c)+".id", gjson.Get(ch.String(), "id").String())
			}
			/*************************************************************************************/
			activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".responses.0.id", "")
			responses := gjson.Get(jsonSR, "surveyresponses.data")
			for rs, res := range responses.Array() {
						activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".responses."+strconv.Itoa(rs)+".id", gjson.Get(res.String(), "id").String())
						//set responses-answers
						curr_que := `pages.0.questions.#[id="`+gjson.Get(que.String(), "id").String()+`"].answers`
						answers := gjson.Get(res.String(), curr_que)
						for a, ans := range answers.Array() {
								activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".responses."+strconv.Itoa(rs)+".answers."+strconv.Itoa(a)+".text", gjson.Get(ans.String(), "text").String())
								activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".responses."+strconv.Itoa(rs)+".answers."+strconv.Itoa(a)+".choice_id", gjson.Get(ans.String(), "choice_id").String())
								activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".responses."+strconv.Itoa(rs)+".answers."+strconv.Itoa(a)+".row_id", gjson.Get(ans.String(), "row_id").String())
								//set answer title from surveydetails
								activityOutput_tmp, _ = sjson.Set(activityOutput_tmp, "survey.questions."+queIndex+".responses."+strconv.Itoa(rs)+".answers."+strconv.Itoa(a)+".title", gjson.Get(que.String(), `answers.choices.#[id="`+gjson.Get(ans.String(), "choice_id").String()+`"].text`).String())
								//fmt.Println("------------>",gjson.Get(ans.String(), "choice_id").String())
						}
			}
			/*************************************************************************************/
			//Update actual output var
			activityOutput = activityOutput_tmp
		}//end of outer loop
		fmt.Println("activityOutput= ",activityOutput)
		return activityOutput
}
