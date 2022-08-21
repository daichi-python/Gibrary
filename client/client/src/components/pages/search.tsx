import React from "react";
import styled from "styled-components";

export const Search: React.FC = () => {
    return (
        <SContainer>
            <h2>Search</h2>
        </SContainer>
    )
}

const SContainer = styled.div`
    display: flex;
    justify-content: center;
    margin-top: 1rem;
`;