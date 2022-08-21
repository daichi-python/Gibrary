import React, { useState } from "react";
import styled from "styled-components";

import { Home } from "../pages/home";
import { Board } from "../pages/board";
import { Search } from "../pages/search";


export const MainTemplate: React.FC = () => {
    const [currentPage, setCurrentPage] = useState("Home");

    const changePage: React.MouseEventHandler<HTMLSpanElement> = (e) => {
        const page = e.currentTarget.innerHTML;
        setCurrentPage(page);
    }

    return (
        <SContainer>
            <SMenuBarContainer>
                <STitleBox>
                    <STitle>Gibrary</STitle>
                </STitleBox>
                <SMenuListBox>
                    <SMenuItemBox>
                        <SMenuItem onClick={changePage}>Home</SMenuItem>
                    </SMenuItemBox>
                    <SMenuItemBox>
                        <SMenuItem onClick={changePage}>Board</SMenuItem>
                    </SMenuItemBox>
                    <SMenuItemBox>
                        <SMenuItem onClick={changePage}>Search</SMenuItem>
                    </SMenuItemBox>
                </SMenuListBox>
            </SMenuBarContainer>
            <SRightContainer>
                {currentPage === "Home" && <Home />}
                {currentPage === "Board" && <Board />}
                {currentPage === "Search" && <Search />}
            </SRightContainer>
        </SContainer>
    )
}

const SContainer = styled.div`
    padding: 2rem;
    display: grid;
    grid-template-columns: 2fr 9fr;
    height: 88vh;
`;

const SMenuBarContainer = styled.div`
    padding: 1rem;
    border: 1px solid #96969649;
`;

const SRightContainer = styled.div`
    padding: 1rem;
`;

const STitleBox = styled.div`
    padding-left: 1rem;
    margin-bottom: 1rem;
    display: flex;
    justify-content: left;
`;

const STitle = styled.h2`
    font-family: Georgia, 'Times New Roman', Times, serif;
    font-size: 1.8rem;
`;

const SMenuListBox = styled.div`
    padding-left: 1.9rem;
`;

const SMenuItemBox = styled.div`
    margin-bottom: 1.9em;
`;

const SMenuItem = styled.span`
    font-size: 1.3rem;
    font-family: Georgia, 'Times New Roman', Times, serif;
    cursor: pointer;
    &:hover {
        transition: 0.1s ease-in-out;
        color: rgb(80, 196, 52);
        font-weight: bold;
    }
`;
