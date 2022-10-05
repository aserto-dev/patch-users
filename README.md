# patch-users

Patch Users (SUPPORT TOOL)

Recreates user which are using an old, none UUID user ID while preserving all other properties.

## Usage

### Check for existance of users with old IDs

```
patch-users --authz-key=${ASERTO_AUTHZ_KEY} --tenant-id=${ASERTO_TENANT_ID}
```

This will return a list of users with old user IDs like this:

```
2022/10/05 13:59:41 TID:<your tenant-id> UID:AB3561054FF503006A66A4DB oldy@acmecorp.com
```

TID reflects the tenant-id containing the user
UID the current, old user ID followed by the email field of the user for identification


### Patching users with old IDs

```
patch-users --authz-key=${ASERTO_AUTHZ_KEY} --tenant-id=${ASERTO_TENANT_ID} --patch
```

When the `--patch` option is added the output will contain a message per user updated, reflecting:

```
2022/10/05 13:59:41 TID:<your tenant-id> UID:AB3561054FF503006A66A4DB oldy@acmecorp.com
2022/10/05 13:59:42 user <old-user-id> recreated with id:<new-user-id>
```
