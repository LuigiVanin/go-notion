export type User = {
    id: number;
    name: string;
    email: string;
    password: string;
    createdAt: string;
    updatedAt: string;
};

export interface ISigninForm {
    email: string;
    password: string;
}

export interface SignupForm extends ISigninForm {
    email: string;
    name: string;
    password: string;
    passwordConfirm: string;
}

export type SigninResponse = {
    token: string;
    user: User;
};
