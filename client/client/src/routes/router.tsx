import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import { Init } from "../components/pages/init";
import { Signup } from "../components/organisms/signup";
import { Signin } from "../components/organisms/signin";
import { MainTemplate } from "../components/templates/maintemplate";


export const MainRouter: React.FC = () => {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<MainTemplate />} />
                <Route path="/auth" element={<Init Signup={Signup} Signin={Signin} />} />
            </Routes>
        </Router>
    )
}