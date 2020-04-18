import React, {useState} from "react"
import "./CodeArea.css"

function CodeArea(props) {
    return (
        <div className={"CodeArea row mt-4"}>
            <div className="col-12 px-0">
                <textarea spellCheck="false" className="code__area py-3" value={props.code} onChange={event => props.setCode(event.target.value)}/>
            </div>
        </div>
    );
}

export default CodeArea