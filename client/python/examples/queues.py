"""
A full example of creating a queue with all options.
"""

import os

import grpc
from armada_client.client import ArmadaClient
from armada_client.permissions import Permissions, Subject


def creating_full_queues_example(client, queue):
    """
    Creates a queue.

    Will skip if the queue already exists.
    """

    subject = Subject(kind="Group", name="group1")
    permissions = Permissions(subjects=[subject], verbs=["cancel", "reprioritize"])

    resource_limits = {"cpu": 1.0, "memory": 1.0}

    queue_req = client.create_queue_request(
        name=queue,
        priority_factor=3.0,
        user_owners=["user1"],
        group_owners=["group1"],
        resource_limits=resource_limits,
        permissions=[permissions],
    )

    # Make sure we handle the queue already existing
    try:
        client.create_queue(queue_req)

    # Handle the error we expect to maybe occur
    except grpc.RpcError as e:
        code = e.code()
        if code == grpc.StatusCode.ALREADY_EXISTS:
            print(f"Queue {queue} already exists")
            client.update_queue(queue_req)
        else:
            raise e


def workflow():
    """
    Starts a workflow, which includes:
        - Creating a queue
        - Creating a jobset
        - Creating a job
    """

    # The queue and job_set_id that will be used for all jobs
    queue = "test-queues"

    # Ensures that the correct channel type is generated
    if DISABLE_SSL:
        channel = grpc.insecure_channel(f"{HOST}:{PORT}")
    else:
        channel_credentials = grpc.ssl_channel_credentials()
        channel = grpc.secure_channel(
            f"{HOST}:{PORT}",
            channel_credentials,
        )

    client = ArmadaClient(channel)

    creating_full_queues_example(client, queue)


if __name__ == "__main__":
    # Note that the form of ARMADA_SERVER should be something like
    # domain.com, localhost, or 0.0.0.0
    DISABLE_SSL = os.environ.get("DISABLE_SSL", False)
    HOST = os.environ.get("ARMADA_SERVER", "localhost")
    PORT = os.environ.get("ARMADA_PORT", "50051")

    workflow()
    print("Completed Workflow")
