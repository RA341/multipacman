import {type Component, createEffect, createSignal, onCleanup, Show} from 'solid-js';

export interface SnackbarMessage {
    id: string;
    message: string;
    type: 'success' | 'error' | 'info' | 'warning';
    duration?: number;
}

interface SnackbarProps {
    message: SnackbarMessage | null;
    onClose: () => void;
}

const Snackbar: Component<SnackbarProps> = (props) => {
    const [isVisible, setIsVisible] = createSignal(false);
    const [isLeaving, setIsLeaving] = createSignal(false);
    let timeoutId: NodeJS.Timeout | null = null;

    createEffect(() => {
        if (props.message) {
            setIsVisible(true);
            setIsLeaving(false);

            // Auto-hide after duration (default 4 seconds)
            const duration = props.message.duration || 4000;
            timeoutId = setTimeout(() => {
                handleClose();
            }, duration);
        } else {
            setIsVisible(false);
        }
    });

    const handleClose = () => {
        setIsLeaving(true);
        setTimeout(() => {
            setIsVisible(false);
            props.onClose();
            setIsLeaving(false);
        }, 300); // Match the transition duration
    };

    const getTypeStyles = (type: string) => {
        switch (type) {
            case 'success':
                return 'bg-green-500 border-green-600';
            case 'error':
                return 'bg-red-500 border-red-600';
            case 'warning':
                return 'bg-yellow-500 border-yellow-600';
            case 'info':
            default:
                return 'bg-blue-500 border-blue-600';
        }
    };

    const getIcon = (type: string) => {
        switch (type) {
            case 'success':
                return (
                    <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd"
                              d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                              clip-rule="evenodd"/>
                    </svg>
                );
            case 'error':
                return (
                    <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd"
                              d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                              clip-rule="evenodd"/>
                    </svg>
                );
            case 'warning':
                return (
                    <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd"
                              d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
                              clip-rule="evenodd"/>
                    </svg>
                );
            case 'info':
            default:
                return (
                    <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd"
                              d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
                              clip-rule="evenodd"/>
                    </svg>
                );
        }
    };

    onCleanup(() => {
        if (timeoutId) {
            clearTimeout(timeoutId);
        }
    });

    return (
        <Show when={isVisible() && props.message}>
            <div class="fixed bottom-4 left-1/2 transform -translate-x-1/2 z-50">
                <div
                    class={`
            max-w-sm w-full shadow-lg rounded-lg border backdrop-blur-md
            transform transition-all duration-300 ease-in-out
            ${getTypeStyles(props.message!.type)}
            ${isLeaving() ? 'translate-y-full opacity-0' : 'translate-y-0 opacity-100'}
          `}
                >
                    <div class="p-4">
                        <div class="flex items-center">
                            <div class="flex-shrink-0 text-white">
                                {getIcon(props.message!.type)}
                            </div>
                            <div class="ml-3 flex-1">
                                <p class="text-sm font-medium text-white">
                                    {props.message!.message}
                                </p>
                            </div>
                            <div class="ml-4 flex-shrink-0 flex">
                                <button
                                    class="inline-flex text-white hover:text-gray-200 focus:outline-none focus:text-gray-200 transition ease-in-out duration-150"
                                    onClick={handleClose}
                                >
                                    <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                                        <path fill-rule="evenodd"
                                              d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                                              clip-rule="evenodd"/>
                                    </svg>
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </Show>
    );
};

export default Snackbar;