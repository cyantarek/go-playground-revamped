import React, {useState} from "react"
import "./Container.css"
import CodeArea from "../CodeArea/CodeArea";
import ResultArea from "../ResultArea/ResultArea";

function Container(props) {
    // example of using state
    // const [message, setMessage] = useState("");
    const {shortCode} = props.match.params;

    return (
        <div className={"Container"}>
            <CodeArea code={props.code} shortCode={shortCode} setCode={props.setCode}/>
            <ResultArea result={props.result} success={props.success}/>
        </div>
    );
}

export default Container