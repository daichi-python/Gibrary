import axios from "axios";

import { API_URL } from "../auth/const";
import { User } from "../types/jsonType";

export const PostUserInformation = (name: string, email: string) => {
    const user: User = {
        name: name,
        email: email
    } 

    const jsonStrings = JSON.stringify(user);

    const url = API_URL + "/user/create";
    axios.post(url, jsonStrings).then((res) => {
        if (res.status === 200) {
            console.log(res.data);
            return res.data
        }
    }).catch((e) => {
        console.log(e);
    })
}