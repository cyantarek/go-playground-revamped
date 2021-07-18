import React, {useEffect, useState} from 'react';
import './App.css';
import "bootstrap/dist/css/bootstrap.min.css"
import Header from "../Header/Header";

import {FormatCodeRequest, RunCodeRequest, ShareCodeRequest, EmptyRequest} from "../../api/playground/playground_service_pb"
import {PlaygroundClient} from "../../api/playground/playground_service_grpc_web_pb"
import {BrowserRouter, Route} from "react-router-dom";
import Container from "../Container/Container";

let playgroundClient = new PlaygroundClient(process.env.REACT_APP_BACKEND_URL, null, null);

function App() {
    const [code, setCode] = useState(initialCode);
    const [result, setResult] = useState("Write code, hit run and have fun :)");
    const [success, setSuccess] = useState(true);
    const [shareLinkCode, setShareLinkCode] = useState("");

    useEffect(() => {
        let req = new EmptyRequest();

        playgroundClient.ping(req, {}, (err, resp) => {
            if (resp == null) {
                if (err.code >= 1000) {
                    setResult(err.message);
                } else {
                    setResult("We can not contact the server. Please wait after some moment");
                }
            }
        })
    }, []);

    const handleRun = () => {
        setResult("Compiling...");

        let req = new RunCodeRequest();
        req.setCode(code);
        req.setLanguage("go");

        playgroundClient.runCode(req, {}, (err, resp) => {
            if (resp == null) {
                if (err.code >= 1000) {
                    setSuccess(false);
                    setResult(err.message);
                } else {
                    setResult("We can not contact the server. Please wait after some moment");
                }
            } else {
                let data = resp.toObject();

                setSuccess(true);
                setResult(data.output + "\n\n" + `Runtime: ${data.runTime}s`)
            }
        })
    };

    const handleShareCode = () => {
        setResult("Getting Shareable Link...");

        let req = new ShareCodeRequest();
        req.setCode(code);
        req.setLanguage("go");

        playgroundClient.shareCode(req, {}, (err, resp) => {
            if (resp == null) {
                if (err.code >= 1000) {
                    setResult(err.message);
                } else {
                    setResult("We can not contact the server. Please wait after some moment");
                }
            } else {
                let data = resp.toObject();

                setResult("");
                setShareLinkCode(`https://localhost:3012/p/${data.shortCode}`)
            }
        })
    };

    const handleCodeFormat = () => {
        setResult("Formatting...");

        let req = new FormatCodeRequest();
        req.setCode(code);
        req.setLanguage("go");

        playgroundClient.formatCode(req, {}, (err, resp) => {
            if (err != null) {
                console.log(err);
                setResult(err.message);
            } else if (resp == null) {
                setResult("We can not contact the server. Please wait after some moment");
            } else {
                let data = resp.toObject();

                console.log(data);

                setCode(data.formattedCode);
                setResult("")
            }
        })
    };

    return (
        <div className="App container-fluid pt-3">
            <BrowserRouter>
                <Header handleRun={handleRun} handleCodeFormat={handleCodeFormat} shareLink={shareLinkCode}
                        handleShareLink={handleShareCode}/>
                        <Route exact path={"/"} render={(props) => <Container {...props} code={code} result={result} setCode={setCode} success={success}/>}/>
                        <Route exact path={"/p/:shortCode"} render={(props) => <Container {...props} code={code} result={result} setCode={setCode} success={success}/>}/>
            </BrowserRouter>
        </div>
    );
}

const initialCode = `package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, playground")
}`;

export default App;
