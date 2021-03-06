#shellcheck disable=SC2034

test_name="deep_migrate_upgrade"
test_container_name="a1-migration.test"
test_upgrades=true
test_upgrade_strategy="none"
test_diagnostics_filters="~skip-for-deep-upgrade"

# The version deployed here has a bug where it's possible for it to
# serve a CLI that is incomplete. We can't go back in time and fix it,
# but we'll detect this and ignore it if it happens
test_detect_broken_cli=true

# The version deployed here has a bug where it's possible for package
# installations to be interrupted by a deployment-service restart,
# leaving the packages in a corrupt state. We can't go back in time
# and fix it, but we try to identify common cases and fix them in the
# tests.
test_detect_broken_packages=true

CURRENT_OLDEST_VERSION=20180629132035
OLD_MANIFEST_DIR="${A2_ROOT_DIR}/components/automate-deployment/testdata/old_manifests/"
DEEP_UPGRADE_PATH="${OLD_MANIFEST_DIR}/${CURRENT_OLDEST_VERSION}.json"

do_build() {
    do_build_default
    sync_a1_migration_data
}

do_deploy() {
    cat $DEEP_UPGRADE_PATH > $test_manifest_path

    local cli_bin="/bin/chef-automate-${CURRENT_OLDEST_VERSION}"

    download_cli "${CURRENT_OLDEST_VERSION}" "${cli_bin}"

    "${cli_bin}" upgrade-from-v1 "$test_config_path" \
        --hartifacts "$test_hartifacts_path" \
        --override-origin "$HAB_ORIGIN" \
        --upgrade-strategy "$test_upgrade_strategy" \
        --manifest-dir "$test_manifest_path" \
        --admin-password chefautomate \
        --skip-preflight \
        --self-test \
        --yes
}
