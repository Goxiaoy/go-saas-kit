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

import { V1TenantInfo } from './v1-tenant-info';

/**
 *
 * @export
 * @interface V1UserTenant
 */
export interface V1UserTenant {
  /**
   *
   * @type {string}
   * @memberof V1UserTenant
   */
  userId?: string;
  /**
   *
   * @type {string}
   * @memberof V1UserTenant
   */
  tenantId?: string;
  /**
   *
   * @type {V1TenantInfo}
   * @memberof V1UserTenant
   */
  tenant?: V1TenantInfo;
  /**
   *
   * @type {boolean}
   * @memberof V1UserTenant
   */
  isHost?: boolean;
}
