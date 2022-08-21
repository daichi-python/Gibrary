import React from "react";
import { Navigate, useLocation } from "react-router-dom";
import { useSetRecoilState } from "recoil";

import { auth } from "../auth/firebase";
import { UserInformation } from "../store/recoil";

type Props = {
    Child: React.ReactNode
}

export const UserProvider: React.FC<Props> = (props) => {
    const { Child } = props;
    const location = useLocation();
    const setUserInfo = useSetRecoilState(UserInformation);

    if (location.pathname.indexOf("/auth") === -1) {
        // Not authentication page.
        const user = auth.currentUser;
        if (user) {
            setUserInfo({
                Name: user.displayName as string,
                Image: user.photoURL as string
            })
        } else {
            return <Navigate to="/auth" />
        }
    }

    return <>{Child}</>
}