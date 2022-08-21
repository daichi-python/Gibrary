import * as firebase from "firebase/app";
import { getAnalytics } from "firebase/analytics";
import { getAuth } from "firebase/auth";
import {
    FIREBASE_APIKEY,
    FIREBASE_APPID,
    FIREBASE_AUTHDOMAIN,
    FIREBASE_DATABASEURL,
    FIREBASE_MEASUREMENTID,
    FIREBASE_MESSAGINGSENDERID,
    FIREBASE_PROJECTID,
    FIREBASE_STORAGEBUCKET
} from "./const";


const firebaseConfig = {
    apiKey: FIREBASE_APIKEY,
    authDomain: FIREBASE_AUTHDOMAIN,
    databaseURL: FIREBASE_DATABASEURL,
    projectId: FIREBASE_PROJECTID,
    storageBucket: FIREBASE_STORAGEBUCKET,
    messagingSenderId: FIREBASE_MESSAGINGSENDERID,
    appId: FIREBASE_APPID,
    measurementId: FIREBASE_MEASUREMENTID,
}
  
const app = firebase.initializeApp(firebaseConfig);
export const analytics = getAnalytics(app);
export const auth = getAuth()
