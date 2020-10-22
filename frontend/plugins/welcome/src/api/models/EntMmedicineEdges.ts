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
    EntPrescription,
    EntPrescriptionFromJSON,
    EntPrescriptionFromJSONTyped,
    EntPrescriptionToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntMmedicineEdges
 */
export interface EntMmedicineEdges {
    /**
     * Prescriptions holds the value of the prescriptions edge.
     * @type {Array<EntPrescription>}
     * @memberof EntMmedicineEdges
     */
    prescriptions?: Array<EntPrescription>;
}

export function EntMmedicineEdgesFromJSON(json: any): EntMmedicineEdges {
    return EntMmedicineEdgesFromJSONTyped(json, false);
}

export function EntMmedicineEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntMmedicineEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'prescriptions': !exists(json, 'prescriptions') ? undefined : ((json['prescriptions'] as Array<any>).map(EntPrescriptionFromJSON)),
    };
}

export function EntMmedicineEdgesToJSON(value?: EntMmedicineEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'prescriptions': value.prescriptions === undefined ? undefined : ((value.prescriptions as Array<any>).map(EntPrescriptionToJSON)),
    };
}


