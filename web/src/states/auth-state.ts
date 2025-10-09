import Cookies from 'js-cookie'
import { create } from 'zustand'
import type {User} from "@/types/user";

const FIREBASE_TOKEN = 'posbe-firebase-token'
const ACCESS_TOKEN = 'posbe-access-token'
const REFRESH_TOKEN = 'posbe-refresh-token'
const USER_PROFILE = 'posbe-user-profile'
const USER_ID = 'posbe-user-id'

interface AuthState {
  auth: {
    firebaseToken: string
    setFirebaseToken: (firebase: string) => void

    accessToken: string
    setAccessToken: (accessToken: string) => void

    refreshToken: string
    setRefreshToken: (refreshToken: string) => void

    user: User | null
    setUser: (user: User | null) => void

    userId: number | null
    setUserId: (userId: number | null) => void

    reset: () => void
  }
}

export const useAuthStore = create<AuthState>()((set) => {
  const firebaseTokenCookieState = Cookies.get(FIREBASE_TOKEN)
  const initFirebaseToken = firebaseTokenCookieState ? JSON.parse(firebaseTokenCookieState) : ''

  const accessTokenCookieState = Cookies.get(ACCESS_TOKEN)
  const initAccessToken = accessTokenCookieState ? JSON.parse(accessTokenCookieState) : ''

  const refreshTokenCookieState = Cookies.get(REFRESH_TOKEN)
  const initRefreshToken = refreshTokenCookieState ? JSON.parse(refreshTokenCookieState) : ''

  const userProfileCookieState = Cookies.get(USER_PROFILE)
  const initUserProfile = userProfileCookieState ? JSON.parse(userProfileCookieState) : ''

  const userIdCookieState = Cookies.get(USER_ID)
  const initUserId = userIdCookieState ? Number(userIdCookieState) : null

  return {
    auth: {
      firebaseToken: initFirebaseToken,
      setFirebaseToken: (firebaseToken) =>
        set((state) => {
          Cookies.set(FIREBASE_TOKEN, JSON.stringify(firebaseToken))
          return { ...state, auth: { ...state.auth, firebaseToken } }
        }),

      accessToken: initAccessToken,
      setAccessToken: (accessToken) =>
        set((state) => {
          Cookies.set(ACCESS_TOKEN, JSON.stringify(accessToken))
          return { ...state, auth: { ...state.auth, accessToken } }
        }),

      refreshToken: initRefreshToken,
      setRefreshToken: (refreshToken) =>
        set((state) => {
          Cookies.set(REFRESH_TOKEN, JSON.stringify(refreshToken))
          return { ...state, auth: { ...state.auth, refreshToken } }
        }),

      userId: initUserId,
      setUserId: (userId) =>  set((state) => {
        Cookies.set(USER_ID, String(userId))
        return { ...state, auth: { ...state.auth, userId } }
      }),

      user: initUserProfile,
      setUser: (user) =>
        set((state) => {
          // {sameSite: 'None', secure: true}
          Cookies.set(USER_PROFILE, JSON.stringify(user))
          return { ...state, auth: { ...state.auth, user } }
        }),

      reset: () =>
        set((state) => {
          Cookies.remove(ACCESS_TOKEN)
          Cookies.remove(REFRESH_TOKEN)
          Cookies.remove(USER_PROFILE)
          Cookies.remove(USER_ID)
          return {
            ...state,
            auth: {
              ...state.auth,
              user: null,
              accessToken: '',
              refreshToken: '',
              userId: null,
            },
          }
        }),
    },
  }
})