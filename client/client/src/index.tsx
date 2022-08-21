import React from "react";
import { render } from "react-dom";
import { RecoilRoot } from "recoil";

import { MainRouter } from "./routes/router";
import "./i18n/configs";

const App: React.FC = () => {
    return (
        <RecoilRoot>
            <MainRouter />
        </RecoilRoot>
    )
}

const root = document.getElementById("root");
render(<App />, root);