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
/**
 * 
 * @export
 * @interface ControllersPrescription
 */
export interface ControllersPrescription {
    /**
     * 
     * @type {string}
     * @memberof ControllersPrescription
     */
    datetime?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersPrescription
     */
    doctor?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPrescription
     */
    patient?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPrescription
     */
    systemmember?: number;
}

export function ControllersPrescriptionFromJSON(json: any): ControllersPrescription {
    return ControllersPrescriptionFromJSONTyped(json, false);
}

export function ControllersPrescriptionFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersPrescription {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'datetime': !exists(json, 'datetime') ? undefined : json['datetime'],
        'doctor': !exists(json, 'doctor') ? undefined : json['doctor'],
        'patient': !exists(json, 'patient') ? undefined : json['patient'],
        'systemmember': !exists(json, 'systemmember') ? undefined : json['systemmember'],
    };
}

export function ControllersPrescriptionToJSON(value?: ControllersPrescription | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'datetime': value.datetime,
        'doctor': value.doctor,
        'patient': value.patient,
        'systemmember': value.systemmember,
    };
}


