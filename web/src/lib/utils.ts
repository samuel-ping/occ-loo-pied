import { API_URL } from "./constants";

interface occupiedResponse {
    occupied: boolean;
    occupiedStartTime?: Date;
}

export async function getOccupied(): Promise<occupiedResponse> {
    const res = await fetch(API_URL);
    if (!res.ok) {
        throw new Error(`Response status: ${res.status}`);
    }

    const json = await res.json();
    return {
        occupied: json.occupied,
        occupiedStartTime: new Date(json.occupiedStartTime)
    };
}

export async function toggleOccupied(occupied: boolean): Promise<occupiedResponse> {
    const res = await fetch(API_URL, {
        method: 'PUT',
        body: JSON.stringify({ occupied: !occupied })
    });

    const json = await res.json();
    return {
        occupied: json.occupied,
        occupiedStartTime: new Date(json.occupiedStartTime)
    }
}

export interface timeSince {
    hours: number;
    minutes: number;
    seconds: number;
}

/**
 * Calculates the amount of time that has passed since thenDate in hours, minutes, and seconds.
 * @param thenDate the date to calculate the time since from
 * @returns the amount of time since the input date in hours, minutes, and seconds
 */
export function timeSince(thenDate: Date): timeSince {
    let now: number = Date.now();
    let then: number = thenDate.getTime();

    let timeSinceMs = now - then;
    let timeSinceDate = new Date(timeSinceMs)
    
    return {
        hours: timeSinceDate.getHours() - 19, // I don't know why but its showing 19 hours more for some reason but I'll figure it out later
        minutes: timeSinceDate.getMinutes(),
        seconds: timeSinceDate.getSeconds()
    }
}