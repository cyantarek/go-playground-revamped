import React, {useEffect, useState} from "react"
import "./Container.css"
import CodeArea from "../CodeArea/CodeArea";
import ResultArea from "../ResultArea/ResultArea";
import {CodeByIDRequest} from "../../api/playground/playground_service_pb";
import {PlaygroundClient} from "../../api/playground/playground_service_grpc_web_pb";

let playgroundClient = new PlaygroundClient(process.env.REACT_APP_BACKEND_URL, null, null);

function Container(props) {
    // example of using state
    // const [message, setMessage] = useState("");
    const {shortCode} = props.match.params;

    useEffect(() => {
        if (shortCode) {
            let req = new CodeByIDRequest()
            req.setId(shortCode)

            playgroundClient.getCodeByShare(req, {}, (err, resp) => {
                let data = resp.toObject();

                props.setCode(data.code)
            })
        }
    });

    return (
        <div className={"Container"}>
            <CodeArea code={props.code} shortCode={shortCode} setCode={props.setCode}/>
            <ResultArea result={props.result} success={props.success}/>
        </div>
    );
}

export default Container
