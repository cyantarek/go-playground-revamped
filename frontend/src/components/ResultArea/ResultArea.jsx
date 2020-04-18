import React from "react"
import "./ResultArea.css"

function ResultArea(props) {
    let content;

    if (props.result === "Compiling..." || props.result === "Formatting...") {
        content = <textarea disabled spellCheck="false" className="result__area compiling" name="" id=""
                            value={props.result}/>
    } else if (props.success === true) {
        content =
            <textarea disabled spellCheck="false" className="result__area success" name="" id="" value={props.result}/>
    } else {
        content =
            <textarea disabled spellCheck="false" className="result__area failed" name="" id="" value={props.result}/>
    }

    return (
        <div className={"ResultArea row mt-2"}>
            <div className="col-12">
                {content}
            </div>
        </div>
    );
}

export default ResultArea