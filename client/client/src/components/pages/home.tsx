import React from "react";
import styled from "styled-components";

import { HomeItem, User } from "../../types/jsonType";
import { HomeContent } from "../organisms/homeContent";


const user: User = {
    id: "1",
    name: "Tominaga Daichi",
    email: "zerukemu0630@gmail.com"
}

const homeItem: HomeItem = {
    id: "2",
    detail: "This is home item. Test Home Item. React is running.",
    is_opened: "TRUE"
}

export const Home: React.FC = () => {
    return (
        <SContainer>
            <SHomeContentBox>
                <HomeContent User={user} HomeItem={homeItem} />
            </SHomeContentBox>
            <></>
        </SContainer>
    )
}

const SContainer = styled.div`
    display: grid;
    grid-template-columns: 3fr 1fr;
    height: 86vh;
`;

const SHomeContentBox = styled.div`
    overflow-y: scroll;
`;