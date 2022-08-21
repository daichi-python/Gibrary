import { 
    Auth, 
    createUserWithEmailAndPassword, 
    FacebookAuthProvider, 
    GoogleAuthProvider, 
    sendEmailVerification, 
    signInWithEmailAndPassword, 
    signInWithPopup
} from "firebase/auth";
import { AuthFormStateType } from "../types/type";

export const FirebaseSignupHandler = (auth: Auth, state: AuthFormStateType) => {
    const { email, password } = state;

    const actionCodeSettings = {
        url: 'http://localhost:3000',
        handleCodeInApp: true,
    };

    createUserWithEmailAndPassword(auth, email, password).then((UserCredential) => {
        sendEmailVerification(UserCredential.user, actionCodeSettings);
        console.log("Complete to send verification email.");
    }).catch((e) => {
        console.log(e);
        return false;
    });

    return true;
}

export const FirebaseLoginHandler = (auth: Auth, state: AuthFormStateType) => {
    const { email, password } = state;
    signInWithEmailAndPassword(auth, email, password)
        .then((UserCredential) => {
            console.log("Success to login.");
            return UserCredential.user;
        }).catch((e) => {
            const errorCode = e.code;
            const errorMessage = e.message;
            console.log(`Failed to login.\nerror code=${errorCode}\nerror message=${errorMessage}`);
            return;
        })
}

export const FirebaseWithGoogle = (auth: Auth) => {
    const provider = new GoogleAuthProvider();
    signInWithPopup(auth, provider)
        .then((UserCredential) => {
            console.log("Success to login with google.");
            return UserCredential.user;
        }).catch((e) => {
            const errorCode = e.code;
            const errorMessage = e.message;
            console.log(`Failed to login with google.\nerror code=${errorCode}\nerror message=${errorMessage}`);
            return;
        })
}

export const FirebaseWithFacebook = (auth: Auth) => {
    const provider = new FacebookAuthProvider();
    signInWithPopup(auth, provider)
        .then((UserCredential) => {
            console.log("Success to login with facebook.");
            return UserCredential.user;
        }).catch((e) => {
            const errorCode = e.code;
            const errorMessage = e.message;
            console.log(`Failed to login with facebook.\nerror code=${errorCode}\nerror message=${errorMessage}`);
            return;
        })
}