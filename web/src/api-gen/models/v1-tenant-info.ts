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

import { BlobBlobFile } from './blob-blob-file';

/**
 *
 * @export
 * @interface V1TenantInfo
 */
export interface V1TenantInfo {
  /**
   *
   * @type {string}
   * @memberof V1TenantInfo
   */
  id?: string;
  /**
   *
   * @type {string}
   * @memberof V1TenantInfo
   */
  name?: string;
  /**
   *
   * @type {string}
   * @memberof V1TenantInfo
   */
  displayName?: string;
  /**
   *
   * @type {string}
   * @memberof V1TenantInfo
   */
  region?: string;
  /**
   *
   * @type {BlobBlobFile}
   * @memberof V1TenantInfo
   */
  logo?: BlobBlobFile;
}
