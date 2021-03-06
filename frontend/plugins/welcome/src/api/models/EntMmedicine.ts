/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntMmedicineEdges,
    EntMmedicineEdgesFromJSON,
    EntMmedicineEdgesFromJSONTyped,
    EntMmedicineEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntMmedicine
 */
export interface EntMmedicine {
    /**
     * MmedicineName holds the value of the "Mmedicine_Name" field.
     * @type {string}
     * @memberof EntMmedicine
     */
    mmedicineName?: string;
    /**
     * 
     * @type {EntMmedicineEdges}
     * @memberof EntMmedicine
     */
    edges?: EntMmedicineEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntMmedicine
     */
    id?: number;
}

export function EntMmedicineFromJSON(json: any): EntMmedicine {
    return EntMmedicineFromJSONTyped(json, false);
}

export function EntMmedicineFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntMmedicine {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'mmedicineName': !exists(json, 'Mmedicine_Name') ? undefined : json['Mmedicine_Name'],
        'edges': !exists(json, 'edges') ? undefined : EntMmedicineEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntMmedicineToJSON(value?: EntMmedicine | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Mmedicine_Name': value.mmedicineName,
        'edges': EntMmedicineEdgesToJSON(value.edges),
        'id': value.id,
    };
}


