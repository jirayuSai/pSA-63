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
    EntDoctor,
    EntDoctorFromJSON,
    EntDoctorFromJSONTyped,
    EntDoctorToJSON,
    EntMmedicine,
    EntMmedicineFromJSON,
    EntMmedicineFromJSONTyped,
    EntMmedicineToJSON,
    EntPatient,
    EntPatientFromJSON,
    EntPatientFromJSONTyped,
    EntPatientToJSON,
    EntSystemmember,
    EntSystemmemberFromJSON,
    EntSystemmemberFromJSONTyped,
    EntSystemmemberToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPrescriptionEdges
 */
export interface EntPrescriptionEdges {
    /**
     * 
     * @type {EntDoctor}
     * @memberof EntPrescriptionEdges
     */
    doctor?: EntDoctor;
    /**
     * 
     * @type {EntMmedicine}
     * @memberof EntPrescriptionEdges
     */
    mmedicine?: EntMmedicine;
    /**
     * 
     * @type {EntPatient}
     * @memberof EntPrescriptionEdges
     */
    patient?: EntPatient;
    /**
     * 
     * @type {EntSystemmember}
     * @memberof EntPrescriptionEdges
     */
    systemmember?: EntSystemmember;
}

export function EntPrescriptionEdgesFromJSON(json: any): EntPrescriptionEdges {
    return EntPrescriptionEdgesFromJSONTyped(json, false);
}

export function EntPrescriptionEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPrescriptionEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'doctor': !exists(json, 'Doctor') ? undefined : EntDoctorFromJSON(json['Doctor']),
        'mmedicine': !exists(json, 'Mmedicine') ? undefined : EntMmedicineFromJSON(json['Mmedicine']),
        'patient': !exists(json, 'Patient') ? undefined : EntPatientFromJSON(json['Patient']),
        'systemmember': !exists(json, 'Systemmember') ? undefined : EntSystemmemberFromJSON(json['Systemmember']),
    };
}

export function EntPrescriptionEdgesToJSON(value?: EntPrescriptionEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'doctor': EntDoctorToJSON(value.doctor),
        'mmedicine': EntMmedicineToJSON(value.mmedicine),
        'patient': EntPatientToJSON(value.patient),
        'systemmember': EntSystemmemberToJSON(value.systemmember),
    };
}


