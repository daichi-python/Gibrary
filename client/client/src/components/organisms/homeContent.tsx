import React from "react";
import styled from "styled-components";

import { HomeItem, User, Groupy } from "../../types/jsonType";
import { ImageAndName } from "../molecules/ImageAndName";
import { HomeActions } from "../molecules/HomeActions";


type Props = {
    User?: User;
    Groupy?: Groupy;
    HomeItem: HomeItem;
}

export const HomeContent: React.FC<Props> = (props) => {
    const { HomeItem, User, Groupy } = props;
    const owner = User === undefined ? Groupy : User;
    var path = "";

    if (User !== undefined) path = "../../../images/users/";
    else if (Groupy !== undefined) path = "../../../images/groupies/"; 


    return (
        <SContainer>
            <ImageAndName Name={owner?.name as string} Image="https://msp.c.yimg.jp/images/v2/FUTi93tXq405grZVGgDqG7EWq3gkaBV51Iq32TvAcGrFGJLYXPsvUcTXV7P3YYThZwu87gLpK7_7hD_Dhrd1QvpmiimTfkYS9Yk60hdIfx3wd4ctZIyCZlH2hSZrO5eh0cJMRCjwGzUJSb2NjfhcmiCB-F5PdVm9BBdSVKnteicVbyoKD0lxwQMCRbuqBNrKgCA4RGaid3itx7yOlGmFKauWD8Jm637V5HDHj2kVN_mGbT-QpWodcgBuEcC3RXkXHlnlwgCoT9nRv30y4Jtej0n7zM1zBKTAt5gGZIeqr3QpR8rbk_MMFYSBI3m5UNKA/20170719172956.png" />
            <SDetailBox>
                <SDetail>{HomeItem.detail}</SDetail>
            </SDetailBox>
            <SActionBox>
                <HomeActions />
            </SActionBox>
        </SContainer>
    )
}

const SContainer = styled.div`
    border: 1px solid #c0c0c067;
`;

const SDetailBox = styled.div`
    display: flex;
    justify-content: center;
`;

const SDetail = styled.p`
    font-size: 1.2rem;
    width: 70%;
`;

const SActionBox = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
`;