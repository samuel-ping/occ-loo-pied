import { OCCUPIED_API_URL, METRICS_API_URL } from "./constants";


interface occupiedResponse {
    occupied: boolean;
    occupiedStartTime?: Date;
}

export async function getOccupied(): Promise<occupiedResponse> {
    const res = await fetch(OCCUPIED_API_URL);
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
    const res = await fetch(OCCUPIED_API_URL, {
        method: 'PUT',
        body: JSON.stringify({ occupied: !occupied })
    });

    const json = await res.json();
    return {
        occupied: json.occupied,
        occupiedStartTime: new Date(json.occupiedStartTime)
    }
}

interface metricsResponse {
    metrics: metric[]
}

export interface metric {
    id: string;
    startTime: Date;
    endTime: Date;
    duration: timeSince;
}

export async function getMetrics(): Promise<metricsResponse> {
    const res = await fetch(METRICS_API_URL)
    if (!res.ok) {
        throw new Error(`Response status: ${res.status}`);
    }

    const json = await res.json();
    return {
        metrics: json.metrics.map((m: any) => ({
            id: m.id,
            startTime: new Date(m.startTime),
            endTime: new Date(m.endTime),
            duration: nanosecondsToTimeSince(m.duration)
        }))
    }
}

export async function deleteMetric(id: string) {
    const url = METRICS_API_URL + "/" + id
    const res = await fetch(url, {
        method:"DELETE"
    })
    if (!res.ok) {
        throw new Error(`Response status: ${res.status}`);
    }
}

export interface timeSince {
    hours: number;
    minutes: number;
    seconds: number;
    milliseconds: number;
}

/**
 * Calculates time elapsed since thenDate in hours, minutes, and seconds.
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
        seconds: timeSinceDate.getSeconds(),
        milliseconds: timeSinceDate.getUTCMilliseconds() // usually isn't used
    }
}

/**
 * Converts time elapsed in nanoseconds to hours, minutes, and seconds.
 * @param duration time elapsed in nanoseconds
 * @returns time elapsed in hours, minutes, and seconds
 */
function nanosecondsToTimeSince(duration: number): timeSince {
    const milliseconds = duration / 1e6;
    const date = new Date(milliseconds);

    return {
        hours: date.getUTCHours(),
        minutes: date.getUTCMinutes(),
        seconds: date.getUTCSeconds(),
        milliseconds: date.getUTCMilliseconds()
    };
}