export type User = {
    id?: string;
    name: string;
    email: string;
    birthday?: string;
    gender?: string;
    is_active?: string;
}

export type Groupy = {
    id?: string;
    name: string;
    introduce?: string;
    max_people?: string;
    group_key?: string;
    is_opened: string;
}

export type BoardItem = {
    id?: string;
    type: string;
    category: string;
    title: string;
    detail?: string;
    applicants?: string;
    is_opened: string;
}

export type HomeItem = {
    id?: string;
    detail: string;
    is_opened: string;
}