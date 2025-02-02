import { OCCUPIED_API_URL, METRICS_API_URL } from './constants';

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
	};
}

interface metricsResponse {
	metrics: metric[];
	pagination: pagination;
}

export interface metric {
	id: string;
	startTime: Date;
	endTime: Date;
	duration: number;
}

export class pagination {
	totalItems: number = 0;
	page: number = 1;
	totalPages: number = 0;
	nextPage?: number = 2;
	prevPage?: number;
}

export async function getMetrics(page: number, itemsPerPage: number): Promise<metricsResponse> {
	const params = new URLSearchParams({
		page: page.toString(),
		itemsPerPage: itemsPerPage.toString()
	});
	const res = await fetch(METRICS_API_URL + '?' + params);
	if (!res.ok) {
		throw new Error(`Response status: ${res.status}`);
	}

	const json = await res.json();

	let paginationDetails = new pagination();
	paginationDetails.totalItems = json.pagination.totalItems;
	paginationDetails.page = json.pagination.page;
	paginationDetails.totalPages = json.pagination.totalPages;
	paginationDetails.nextPage = json.pagination.nextPage;
	paginationDetails.prevPage = json.pagination.prevPage;

	return {
		metrics: json.metrics.map((m: any) => ({
			id: m.id,
			startTime: new Date(m.startTime),
			endTime: new Date(m.endTime),
			duration: m.duration
		})),
		pagination: paginationDetails
	};
}

export async function deleteMetric(id: string) {
	const url = METRICS_API_URL + '/' + id;
	const res = await fetch(url, {
		method: 'DELETE'
	});
	if (!res.ok) {
		throw new Error(`Response status: ${res.status}`);
	}
}

interface usagesByDayResponse {
	usagesByDay: usageByDayMetric[];
	mostUsagesInADay: number;
}

export interface usageByDayMetric {
	date: string;
	timesUsed: number;
}

export async function usagesByDay(): Promise<usagesByDayResponse> {
	const res = await fetch(METRICS_API_URL + '/usagesByDay');
	if (!res.ok) {
		throw new Error(`Response status: ${res.status}`);
	}

	const json = await res.json();

	return {
		usagesByDay: json.usagesByDay.map((u: any) => ({
			date: u.date,
			timesUsed: u.timesUsed
		})),
		mostUsagesInADay: json.mostUsagesInADay
	}
}

export class stats {
	totalUsages: number = 0;
	duration: durationStats = {
		total: -1,
		longest: {
			id: "",
			duration: -1,
			startTime: new Date(),
			endTime: new Date(),
		},
		average: -1.0
	};
}

export interface durationStats {
	total: number;
	longest: metric;
	average: number;
}

export async function getStats(): Promise<stats> {
	const res = await fetch(METRICS_API_URL + '/stats');
	if (!res.ok) {
		throw new Error(`Response status: ${res.status}`);
	}

	const json = await res.json();

	json.stats.duration.longest.startTime = new Date(json.stats.duration.longest.startTime);
	json.stats.duration.longest.endTime = new Date(json.stats.duration.longest.endTime);
	
	return json.stats
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
	let timeSinceDate = new Date(timeSinceMs);

	return {
		hours: timeSinceDate.getUTCHours(),
		minutes: timeSinceDate.getMinutes(),
		seconds: timeSinceDate.getSeconds(),
		milliseconds: timeSinceDate.getUTCMilliseconds() // usually isn't used
	};
}

/**
 * Converts time elapsed in nanoseconds to hours, minutes, and seconds.
 * @param duration time elapsed in nanoseconds
 * @returns time elapsed in hours, minutes, and seconds
 */
export function nanosecondsToTimeSince(duration: number): timeSince {
	const milliseconds = duration / 1e6;
	const date = new Date(milliseconds);

	return {
		hours: date.getUTCHours(),
		minutes: date.getMinutes(),
		seconds: date.getSeconds(),
		milliseconds: date.getMilliseconds()
	};
}
