/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/mule
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

import "encoding/base64"

// LicenseResource is a generated function returning the Resource as a []byte.
func ReadmeResource() ([]byte, error) {
	var resource = "DQpFbmQgVXNlciBMaWNlbnNlIEFncmVlbWVudA0KDQpUaGlzIGxpY2Vuc2UgYWdyZWVtZW50IGlzIGEgbGVnYWwgYWdyZWVtZW50IGJldHdlZW4geW91LCB0aGUgImVuZCB1c2VyIiAoZWl0aGVyIGFzIGFuIGluZGl2aWR1YWwgb3IgYW4gZW50aXR5KSwgYW5kIFBpbmdtYW4gVG9vbHMuDQoNClRoZSBzb2Z0d2FyZSBhbmQgYW55IG9mIGl0cyBzdXBwb3J0aW5nIGRvY3VtZW50YXRpb24gYWZ0ZXIgcmVmZXJyZWQgdG8gYXMgdGhlICJwcm9kdWN0Ii4gQnkgaW5zdGFsbGluZyBvciB1c2luZyB0aGUgcHJvZHVjdCwgeW91IGluZGljYXRlIHlvdXIgY29tcGxldGUgYW5kIHVuY29uZGl0aW9uYWwgYWNjZXB0YW5jZSBvZiB0aGVzZSB0ZXJtcyBhbmQgY29uZGl0aW9ucy4gSWYgeW91IGRvIG5vdCBhZ3JlZSB0byB0aGUgdGVybXMgYW5kIGNvbmRpdGlvbnMgb2YgdGhpcyBhZ3JlZW1lbnQgcHJvbXB0bHkgcmVtb3ZlIHRoZSBwcm9kdWN0IGZvcm0geW91ciBjb21wdXRlciBhbmQgZGVzdHJveSBhc3NvY2lhdGVkIGRvY3VtZW50YXRpb24uDQoNClNob3VsZCB0aGVyZSBiZSBhbnkgY29uZmxpY3QgYmV0d2VlbiB0aGUgdGVybXMgYW5kIGNvbmRpdGlvbnMgb2YgdGhpcyBhZ3JlZW1lbnQgYW5kIHRoZSB0ZXJtcyBhbmQgY29uZGl0aW9ucyBvZiBhbnkgb3RoZXIgYWdyZWVtZW50IGJldHdlZW4geW91IGFuZCBQaW5nbWFuIFRvb2xzIG9yIHRoZWlyIHNlcnZhbnRzIG9yIGFnZW50cyBpbiByZWxhdGlvbiB0byB0aGUgcHJvZHVjdCwgdGhlIHRlcm1zIGFuZCBjb25kaXRpb25zIG9mIHRoaXMgYWdyZWVtZW50IHNoYWxsIGFwcGx5Lg0KDQpJZiBhbnkgcHJvdmlzaW9uIG9mIHRoaXMgYWdyZWVtZW50IGlzIGZvdW5kIHRvIGJlIHVubGF3ZnVsIG9yIHZvaWQgdGhlbiB0aGF0IHByb3Zpc2lvbiBzaGFsbCBiZSBzZXZlcmVkIGZyb20gdGhpcyBhZ3JlZW1lbnQgYW5kIHdpbGwgbm90IGFmZmVjdCB0aGUgdmFsaWRpdHkgb2YgdGhlIHJlbWFpbmluZyBwcm92aXNpb25zLg0KDQpJbmZvcm1hdGlvbiBjb250YWluZWQgaW4gdGhpcyBhZ3JlZW1lbnQgaXMgc3ViamVjdCB0byBjaGFuZ2Ugd2l0aG91dCBub3RpY2UgYW5kIGRvZXMgbm90IHJlcHJlc2VudCBhIGNvbW1pdG1lbnQgb3IgY29udHJhY3RlZCBvYmxpZ2F0aW9uIG9uIHRoZSBwYXJ0IG9mIFBpbmdtYW4gVG9vbHMuDQpDb3B5cmlnaHQgTm90aWNlDQoNCihDKSBDb3B5cmlnaHQgMTk5OCwgMjAxNiBQaW5nbWFuIFRvb2xzLiBBbGwgcmlnaHRzIHJlc2VydmVkLg0KDQpUaGUgcHJvZHVjdCBpcyBvd25lZCBieSBQaW5nbWFuIFRvb2xzLiBDb3B5cmlnaHQgbGF3cyBhbmQgaW50ZXJuYXRpb25hbCBjb3B5cmlnaHQgdHJlYXRpZXMsIGFzIHdlbGwgYXMgb3RoZXIgaW50ZWxsZWN0dWFsIHByb3BlcnR5IGxhd3MgYW5kIHRyZWF0aWVzIHByb3RlY3QgdGhlIFByb2R1Y3QuIFRoZSBQcm9kdWN0IGlzIGxpY2Vuc2VkLCBub3Qgc29sZC4gQWxsIHJpZ2h0cyBpbiB0aGUgcHJvZHVjdCBpbmNsdWRpbmcgTGljZW5zZSBBZ3JlZW1lbnQsIGNvcHlyaWdodHMsIGxpY2Vuc2luZyByaWdodHMsIHBhdGVudHMsIHRyYWRlbWFya3MsIHRyYWRlIHNlY3JldHMsIGRlc2lnbiByaWdodHMsIGVuZ2luZWVyaW5nIHJpZ2h0cywgbW9yYWwgcmlnaHRzLCBhbmQgYW55IG90aGVyIGludGVsbGVjdHVhbCBwcm9wZXJ0eSByaWdodHMgYmVsb25nIHRvIFBpbmdtYW4gVG9vbHMuIFRoZXNlIHJpZ2h0cyBhcmUgbm90IHRyYW5zZmVycmVkIGFzIHBhcnQgb2YgdGhpcyBhZ3JlZW1lbnQuDQoNCk5vIHBhcnQgb2YgdGhlIHByb2R1Y3QsIGluY2x1ZGluZyB0aGUgTGljZW5zZSBLZXkocyksIG1heSBiZSByZXByb2R1Y2VkLCBwdWJsaXNoZWQsIHRyYW5zbWl0dGVkIGVsZWN0cm9uaWNhbGx5LCBtZWNoYW5pY2FsbHkgb3Igb3RoZXJ3aXNlLCB0cmFuc2NyaWJlZCwgc3RvcmVkIGluIGEgcmV0cmlldmFsIHN5c3RlbSBvciB0cmFuc2xhdGVkIGludG8gYW55IGxhbmd1YWdlIGluIGFueSBmb3JtLCBieSBhbnkgbWVhbnMsIGZvciBhbnkgcHVycG9zZSBvdGhlciB0aGFuIHRoZSBwdXJjaGFzZXIncyBwZXJzb25hbCB1c2UsIHdpdGhvdXQgdGhlIGV4cHJlc3Mgd3JpdHRlbiBwZXJtaXNzaW9uIG9mIFBpbmdtYW4gVG9vbHMuDQpFdmFsdWF0aW9uIExpY2Vuc2UgQWdyZWVtZW50DQoNClRoaXMgRXZhbHVhdGlvbiBMaWNlbnNlIEFncmVlbWVudCBncmFudHMgeW91IGEgbm9uLWV4Y2x1c2l2ZSBMaWNlbnNlIHRvIHVzZSB0aGUgcHJvZHVjdCBmb3IgYW4gZXZhbHVhdGlvbiBwZXJpb2Qgb2YgZm91cnRlZW4gKDE0KSBkYXlzIGZyb20gdGhlIGRhdGUgeW91IGluc3RhbGwgdGhlIHByb2R1Y3QuDQoNCk9uIHRoZSBmaWZ0ZWVudGggKDE1KSBkYXkgYWZ0ZXIgeW91IGluc3RhbGxlZCB0aGUgcHJvZHVjdCB5b3UgbXVzdCBlaXRoZXIgcmVnaXN0ZXIgdGhlIHByb2R1Y3QgYnkgbWVhbnMgb2YgcHVyY2hhc2luZyBhIGxpY2Vuc2UgKHNpbmdsZSB1c2VyLCBtdWx0aSBjb21wdXRlciBvciBhIHNpdGUgbGljZW5zZSkgZnJvbSBQaW5nbWFuIFRvb2xzLCBvciBkZXN0cm95IGFsbCBjb3BpZXMgb2YgdGhlIHByb2R1Y3QgaW4geW91ciBwb3NzZXNzaW9uLCBhbmQgYW55IHJlbGF0ZWQgZG9jdW1lbnRhdGlvbi4NClNvZnR3YXJlIFVzYWdlDQoNClVzZSBEZWZpbml0aW9uLiBUaGUgcHJvZHVjdCBpcyAiaW4gdXNlIiBvbiBhIGNvbXB1dGVyIHdoZW4gaXQgaXMgbG9hZGVkIGludG8gdGVtcG9yYXJ5IG1lbW9yeSAoaS5lLiBSQU0pIG9yIGluc3RhbGxlZCBpbnRvIHBlcm1hbmVudCBtZW1vcnkgKGUuZy4gaGFyZCBkaXNrLCBDRC1ST00sIG9yIG90aGVyIHN0b3JhZ2UgZGV2aWNlKSBvZiB0aGF0IGNvbXB1dGVyLCBleGNlcHQgdGhhdCBhIGNvcHkgaW5zdGFsbGVkIG9uIGEgbmV0d29yayBzZXJ2ZXIgZm9yIHRoZSBzb2xlIHB1cnBvc2Ugb2YgZGlzdHJpYnV0aW9uIHRvIG90aGVyIGNvbXB1dGVycyBpcyBub3QgImluIHVzZSIuDQpTaW5nbGUgdXNlciBsaWNlbnNlDQoNCkEgc2luZ2xlIHVzZXIgbGljZW5zZSBpcyBhIGxpY2Vuc2UgdG8gYSBzaW5nbGUgdXNlci4gVGhpcyBzaW5nbGUgdXNlciBjYW4gdXNlIHRoZSBwcm9kdWN0IG9uIGFueSBjb21wdXRlciBhcyBsb25nIGFzIHRoZSBwcm9kdWN0IGlzIG5vdCBpbiBjb25jdXJyZW50IHVzZSBvbiBhbnkgb3RoZXIgY29tcHV0ZXIuIGkuZS4geW91IG93biB0d28gY29tcHV0ZXJzLiBZb3UgbWF5IGhhdmUgdGhlIHByb2R1Y3QgaW5zdGFsbGVkLCBidXQgbm90IGhhdmUgaXQgcnVubmluZyBvbiBtb3JlIHRoYW4gb25lIGNvbXB1dGVyIGF0IGEgdGltZS4gVGhpcyAicmVhc29uYWJsZSB1c2UiIGNsYXVzZSBhcHBsaWVzIG9ubHkgdG8gdGhlIHNpbmdsZSB1c2VyIGxpY2Vuc2UuDQpNdWx0aSBjb21wdXRlciBsaWNlbnNlDQoNCllvdSBtYXkgaW5zdGFsbCB0aGUgcHJvZHVjdGlvbiB2ZXJzaW9uIG9mIHRoZSBwcm9kdWN0IHRvIGFueSBudW1iZXIgb2YgY29tcHV0ZXJzIHByb3ZpZGVkIHRoYXQgdGhpcyBudW1iZXIgZG9lcyBub3QgZXhjZWVkIHRoZSBsaW1pdCBzZXQgaW4gdGhlIGxpY2Vuc2UgZmlsZSBwdXJjaGFzZWQgYnkgeW91LiBJZiB5b3Ugb3duIGEgMiBjb21wdXRlciBsaWNlbnNlLCB5b3UgbWF5IG5vdCBpbnN0YWxsIHRoZSBwcm9kdWN0IG9uIDMgY29tcHV0ZXJzLCBldmVuIGlmIHlvdSBhcmUgb25seSB1c2luZyBpdCBvbiAyIGNvbXB1dGVycyBjb25jdXJyZW50bHkuDQpTaW5nbGUgc2l0ZSBsaWNlbnNlDQoNCkEgc2luZ2xlIHNpdGUgbGljZW5zZSBhdXRob3JpemVzIHlvdSB0byBpbnN0YWxsIGFuZCB1c2UgdGhlIHByb2R1Y3QgdG8gYW55IG51bWJlciBvZiBjb21wdXRlcnMgd2l0aGluIGEgc2luZ2xlIHNpdGUuDQpNdWx0aSBzaXRlIGxpY2Vuc2UNCg0KQSBtdWx0aSBzaXRlIGxpY2Vuc2UgYXV0aG9yaXplcyB5b3UgdG8gaW5zdGFsbCBhbmQgdXNlIHRoZSBwcm9kdWN0IHRvIGFueSBudW1iZXIgb2YgY29tcHV0ZXJzIGJlbG9uZ2luZyB0byB5b3VyIG9yZ2FuaXphdGlvbiAtIG5vIG1hdHRlciB3aGVyZSB0aGV5IGFyZSBsb2NhdGVkLg0KTGltaXRhdGlvbiBvbiBVc2UNCg0KWW91IG1heSBub3Q6IHBlcm1pdCBvdGhlciBpbmRpdmlkdWFscyB0byB1c2UgdGhlIHByb2R1Y3QgZXhjZXB0IHVuZGVyIHRoZSB0ZXJtcyBsaXN0ZWQgYWJvdmU7IHBlcm1pdCBjb25jdXJyZW50IHVzZSBvZiB0aGUgcHJvZHVjdCBvdGhlciB0aGFuIHNwZWNpZmllZCB1bmRlciB0aGUgU29mdHdhcmUgVXNhZ2Ugc2VjdGlvbiBvZiB0aGlzIExpY2Vuc2U7IG1vZGlmeSwgdHJhbnNsYXRlLCByZXZlcnNlIGVuZ2luZWVyLCBkZWNvbXBpbGUsIGRlY3J5cHQsIGV4dHJhY3QsIGRpc2Fzc2VtYmxlLCBvciBjcmVhdGUgZGVyaXZhdGl2ZSB3b3JrcyBiYXNlZCBvbiB0aGUgcHJvZHVjdDsgY29weSB0aGUgcHJvZHVjdCBvdGhlciB0aGFuIGFzIHNwZWNpZmllZCBpbiBTb2Z0d2FyZSBVc2FnZSBzZWN0aW9uIG9mIHRoaXMgTGljZW5zZTsgc2VsbCwgcmVudCwgbGVhc2UsIGdyYW50IGEgc2VjdXJpdHkgaW50ZXJlc3QgaW4sIG9yIG90aGVyd2lzZSB0cmFuc2ZlciByaWdodHMgdG8gdGhlIHByb2R1Y3Q7IG9yIGFsdGVyIG9yIHJlbW92ZSBhbnkgcHJvcHJpZXRhcnkgbm90aWNlcyBvciBsYWJlbHMgb24gdGhlIHByb2R1Y3QuIExpY2Vuc2VlIHdhcnJhbnRzIHRoYXQgaXQgd2lsbCBub3QgdXNlIG9yIHJlZGlzdHJpYnV0ZSB0aGUgcHJvZHVjdCBmb3Igc3VjaCBwdXJwb3Nlcy4NCkxpbWl0YXRpb24gb2YgTGlhYmlsaXR5DQoNClBpbmdtYW4gVG9vbHMgc2hhbGwgbm90IGluIGFueSBldmVudCBiZSBsaWFibGUgZm9yIGRpcmVjdCwgaW5kaXJlY3QsIHNwZWNpYWwsIGluY2lkZW50YWwgb3IgY29uc2VxdWVudGlhbCBkYW1hZ2VzIHJlc3VsdGluZyBmcm9tIGFueSBkZWZlY3QgaW4gdGhlIHByb2R1Y3QgaW5jbHVkaW5nIGRhbWFnZXMgZnJvbSBsb3NzIG9mIGRhdGEsIGRvd250aW1lLCBnb29kd2lsbCwgZGFtYWdlIHRvIG9yIHJlcGxhY2VtZW50IG9mIHNvZnR3YXJlLCBlcXVpcG1lbnQgb3IgcHJvcGVydHksIG9yIGFueSBjb3N0cyBvZiByZWNvdmVyaW5nLCByZS1wcm9ncmFtbWluZywgb3IgcmVwcm9kdWNpbmcgYW55IHByb2dyYW0gb3IgZGF0YSB1c2VkIGluIGNvbmp1bmN0aW9uIHdpdGggUGluZ21hbiBUb29scyBvciB0aGUgcHJvZHVjdCwgZXZlbiBpZiBQaW5nbWFuIFRvb2xzIG9yIGFuIGF1dGhvcml6ZWQgZGVhbGVyIGhhcyBiZWVuIGFkdmlzZWQgb2YgdGhlIHBvc3NpYmlsaXR5IG9mIHN1Y2ggZGFtYWdlcy4NCg0KU29tZSBzdGF0ZXMgZG8gbm90IGFsbG93IHRoZSBleGNsdXNpb24gb3IgbGltaXRhdGlvbiBvZiBpbXBsaWVkIHdhcnJhbnRpZXMgb3IgbGlhYmlsaXR5IGZvciBpbmNpZGVudGFsIG9yIGNvbnNlcXVlbnRpYWwgZGFtYWdlcywgc28gdGhlIGFib3ZlIGxpbWl0YXRpb24gbWF5IG5vdCBhcHBseSB0byB5b3UuIFRoaXMgTGltaXRhdGlvbiBvZiBMaWFiaWxpdHkgZ2l2ZXMgeW91IHNwZWNpZmljIGxlZ2FsIHJpZ2h0cywgYW5kIHlvdSBtYXkgYWxzbyBoYXZlIG90aGVyIHJpZ2h0cywgd2hpY2ggdmFyeSwgZnJvbSBzdGF0ZSB0byBzdGF0ZSBvciBjb3VudHJ5IHRvIGNvdW50cnkuDQpDdXN0b21lcidzIExpYWJpbGl0eQ0KDQpUaGUgTGljZW5zZSBLZXkgc3VwcGxpZWQgYXMgcGFydCBvZiB0aGUgcHJvZHVjdCByZW1haW5zIHRoZSBwcm9wZXJ0eSBhbmQgY29weXJpZ2h0IG9mIFBpbmdtYW4gVG9vbHMuIEl0IGlzIGEgYnJlYWNoIG9mIHRoaXMgTGljZW5zZSBBZ3JlZW1lbnQgdG8gc2hhcmUgdGhlIExpY2Vuc2UgS2V5IHdpdGggYW55b25lIG9yIGFueSBlbnRpdHkgZm9yIGFueSBwdXJwb3NlIG90aGVyIHRoYW4gdGhlIHB1cmNoYXNlcidzIHVzZSAtIHdpdGhpbiB0aGUgY29uZGl0aW9ucyBvZiB0aGUgU29mdHdhcmUgVXNhZ2Ugc2VjdGlvbi4gWW91IHdpbGwgYmUgaGVsZCBsaWFibGUgZm9yIG1pc3VzZSBvciBwdWJsaWNhdGlvbiAoaW4gYW55IGZvcm0pIG9mIHRoZSBSZWdpc3RyYXRpb24gQ29kZS4NCldhcnJhbnR5DQoNClBpbmdtYW4gVG9vbHMgd2FycmFudHMgdGhhdCAoYSkgdGhlIHByb2R1Y3Qgd2lsbCBwZXJmb3JtIHN1YnN0YW50aWFsbHkgaW4gYWNjb3JkYW5jZSB3aXRoIHRoZSBhY2NvbXBhbnlpbmcgZG9jdW1lbnRhdGlvbiBmb3IgYSBwZXJpb2Qgb2YgbmluZXR5ICg5MCkgZGF5cyBmcm9tIHRoZSBkYXRlIG9mIHJlY2VpcHQsIGFuZCAoYikgUGluZ21hbiBUb29scyBzdXBwb3J0IGVuZ2luZWVycyB3aWxsIG1ha2UgY29tbWVyY2lhbGx5IHJlYXNvbmFibGUgZWZmb3J0cyB0byBzb2x2ZSBhbnkgcHJvYmxlbSBpc3N1ZXMuIFNvbWUgc3RhdGVzIGFuZCBqdXJpc2RpY3Rpb25zIGRvIG5vdCBhbGxvdyBsaW1pdGF0aW9ucyBvbiBkdXJhdGlvbiBvZiBhbiBpbXBsaWVkIHdhcnJhbnR5LCBzbyB0aGUgYWJvdmUgbGltaXRhdGlvbiBtYXkgbm90IGFwcGx5IHRvIHlvdS4gVG8gdGhlIGV4dGVudCBhbGxvd2VkIGJ5IGFwcGxpY2FibGUgbGF3LCBpbXBsaWVkIHdhcnJhbnRpZXMgb24gdGhlIHByb2R1Y3QsIGlmIGFueSwgYXJlIGxpbWl0ZWQgdG8gbmluZXR5ICg5MCkgZGF5cy4NCkV4cG9ydCBSZWd1bGF0aW9ucw0KDQpTb2Z0d2FyZSwgaW5jbHVkaW5nIHRlY2huaWNhbCBkYXRhLCBpcyBzdWJqZWN0IHRvIFUuUy4gZXhwb3J0IGNvbnRyb2wgbGF3cywgaW5jbHVkaW5nIHRoZSBVLlMuIEV4cG9ydCBBZG1pbmlzdHJhdGlvbiBBY3QgYW5kIGl0cyBhc3NvY2lhdGVkIHJlZ3VsYXRpb25zLCBhbmQgbWF5IGJlIHN1YmplY3QgdG8gZXhwb3J0IG9yIGltcG9ydCByZWd1bGF0aW9ucyBpbiBvdGhlciBjb3VudHJpZXMuIExpY2Vuc2VlIGFncmVlcyB0byBjb21wbHkgc3RyaWN0bHkgd2l0aCBhbGwgc3VjaCByZWd1bGF0aW9ucyBhbmQgYWNrbm93bGVkZ2VzIHRoYXQgaXQgaGFzIHRoZSByZXNwb25zaWJpbGl0eSB0byBvYnRhaW4gbGljZW5zZXMgdG8gZXhwb3J0LCByZS1leHBvcnQsIG9yIGltcG9ydCBTb2Z0d2FyZS4gU29mdHdhcmUgbWF5IG5vdCBiZSBkb3dubG9hZGVkLCBvciBvdGhlcndpc2UgZXhwb3J0ZWQgb3IgcmUtZXhwb3J0ZWQgKGkpIGludG8sIG9yIHRvIGEgbmF0aW9uYWwgb3IgcmVzaWRlbnQgb2YgYW55IGNvdW50cnkgdG8gd2hpY2ggdGhlIFUuUy4gaGFzIGVtYmFyZ29lZCBnb29kczsgb3IgKGlpKSB0byBhbnlvbmUgb24gdGhlIFUuUy4gVHJlYXN1cnkgRGVwYXJ0bWVudCdzIGxpc3Qgb2YgU3BlY2lhbGx5IERlc2lnbmF0ZWQgTmF0aW9ucyBvciB0aGUgVS5TLiBDb21tZXJjZSBEZXBhcnRtZW50J3MgVGFibGUgb2YgRGVuaWFsIE9yZGVycy4NCg=="

	return base64.StdEncoding.DecodeString(resource)
}