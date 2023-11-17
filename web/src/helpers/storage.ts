import { User } from "@/types/user.ts";

interface StorageProxy<T> {
    getItem(): T | null;
    setItem(value: T | null): void;
    removeItem(): void;
}

export class Storage<T> implements StorageProxy<T> {
    constructor(private key: string) {}

    getItem(): T | null {
        const result = localStorage.getItem(this.key);

        if (!result) return null;

        try {
            return JSON.parse(result) as T;
        } catch {
            return null;
        }
    }

    setItem(value: T | null): void {
        const serialized = JSON.stringify(value);
        localStorage.setItem(this.key, serialized);
    }

    removeItem(): void {
        localStorage.removeItem(this.key);
    }
}

export const USER_STORAGE_KEY = "writtable-user";
export const AUTH_STORAGE_KEY = "writtable-auth";

export const userStorage = new Storage<User>(USER_STORAGE_KEY);
export const authStorage = new Storage<{ jwt: string }>(AUTH_STORAGE_KEY);
