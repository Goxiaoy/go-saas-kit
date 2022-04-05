/* tslint:disable */
/* eslint-disable */
/**
 * Saas Service
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { V1TenantConnectionString } from './v1-tenant-connection-string';
import { V1TenantFeature } from './v1-tenant-feature';

/**
 *
 * @export
 * @interface V1UpdateTenant
 */
export interface V1UpdateTenant {
  /**
   *
   * @type {string}
   * @memberof V1UpdateTenant
   */
  id: string;
  /**
   *
   * @type {string}
   * @memberof V1UpdateTenant
   */
  name: string;
  /**
   *
   * @type {string}
   * @memberof V1UpdateTenant
   */
  displayName?: string;
  /**
   *
   * @type {Array<V1TenantConnectionString>}
   * @memberof V1UpdateTenant
   */
  conn?: Array<V1TenantConnectionString>;
  /**
   *
   * @type {Array<V1TenantFeature>}
   * @memberof V1UpdateTenant
   */
  features?: Array<V1TenantFeature>;
  /**
   *
   * @type {string}
   * @memberof V1UpdateTenant
   */
  logo?: string;
}
