import React, {useState} from "react"
import "./Header.css"

function Header(props) {
    return (
        <div className={"Header row"}>
            <div className="col-10">
                <div className="row">
                    <h2 className="logo mx-1">The Go Playground</h2>
                    <button className="btn btn__main mx-1" onClick={props.handleRun}>Run</button>
                    <button className="btn btn__main mx-1" onClick={props.handleCodeFormat}>Format</button>
                    <button className="btn btn__main mx-1" onClick={props.handleShareLink}>Share</button>
                    {props.shareLink ? (<input spellCheck="false" type="email" readOnly={true} className="input__share mx-1"
                        value={props.shareLink}/>) : null}
                </div>
            </div>
            <div className="col-2 d-flex justify-content-end">
                <button className="btn btn__main mx-1">About</button>
            </div>
        </div>
    );
}

export default Header